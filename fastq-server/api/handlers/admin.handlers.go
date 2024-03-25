package handlers

import (
	"context"
	"fmt"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"sync"

	"github.com/bentenison/fastq-server/api/apperrors"
	"github.com/bentenison/fastq-server/models"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}
var clients = make(map[*websocket.Conn]string)
var clientsMutex sync.Mutex

type WSpayload struct {
	IP            string        `json:"ip,omitempty"`
	Action        string        `json:"action,omitempty"`
	TicketPayload models.Ticket `json:"ticket_payload,omitempty"`
}

var broadcast = make(chan WSpayload)

func (h *Handler) wsHandler(c *gin.Context) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	// upgrade this connection to a WebSocket
	// connection
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
	}
	clientIp := c.ClientIP()
	h.AdminService.SetActiveService(clientIp, true)
	// helpful log statement to show connections
	log.Println("Client Connected")
	clientsMutex.Lock()
	clients[conn] = clientIp
	clientsMutex.Unlock()

	defer func() {
		// Remove client from the map when connection closes
		clientsMutex.Lock()
		delete(clients, conn)
		clientsMutex.Unlock()
		conn.Close()
		h.AdminService.SetActiveService(clientIp, false)
	}()

	for {
		// Read message from the client
		msg := WSpayload{}
		err := conn.ReadJSON(&msg)
		if err != nil {
			log.Println("error in reading message", err)
			clientsMutex.Lock()
			delete(clients, conn)
			clientsMutex.Unlock()
			conn.Close()
			h.AdminService.SetActiveService(clientIp, false)
			return
		}
		log.Println("string", msg)

		// Process the message (execute corresponding function)
		// response := h.handleMessage(conn, p, clientIp)
		broadcast <- msg
		// Send the response back to the client

	}
}
func (h *Handler) getConnectedClients(c *gin.Context) {
	// Retrieve and return the list of currently connected clients
	clientsMutex.Lock()
	defer clientsMutex.Unlock()

	connectedClients := make(map[string]string)
	for _, id := range clients {
		connectedClients[id] = "Connected"
	}

	c.JSON(http.StatusOK, connectedClients)
}

// func processMessage(request string) string {
// 	// Implement your server-side logic here based on the client request
// 	// For example, execute functions and return data

//		// This is a simplified example echoing the received request
//		return "Server Response: " + request
//	}
func handleMessage() {
	// Implement your message handling logic here
	// Example: Check if the message indicates the client is disconnecting
	// payload := WSpayload{}
	// log.Println("websocket message recieved from client", string(message))
	// if string(message) == "disconnect" {
	// 	clientsMutex.Lock()
	// 	delete(clients, conn)
	// 	clientsMutex.Unlock()
	// 	log.Printf("Client %s disconnected.\n", clientID)
	// 	return "disconnected"
	// }
	// err := json.Unmarshal(message, &payload)
	// if err != nil {
	// 	log.Println("error occured while unmarshaling message", err)
	// 	return ""
	// }
	// dsdetails, err := h.TicketService.GetSystemService(context.TODO())
	// if err != nil {
	// 	log.Println("error occured while unmarshaling message", err)
	// 	return ""
	// }
	// payload.IP = dsdetails.ServerIP
	// log.Println(payload)
	// res, err := json.Marshal(payload)
	// if err != nil {
	// 	log.Println("error occured while marshaling message", err)
	// 	return ""
	// }
	// if err := conn.WriteMessage(messageType, p); err != nil {
	// 	log.Println("error in writing message", err)
	// 	return []byte{}
	// }
	// return string(res)
	for {
		msg := <-broadcast
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("Error: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}

// func reader(conn *websocket.Conn) {
// 	for {
// 		// read in a message
// 		messageType, p, err := conn.ReadMessage()
// 		if err != nil {
// 			log.Println(err)
// 			return
// 		}
// 		// print out that message for clarity
// 		log.Println(string(p))

// 		if err := conn.WriteMessage(messageType, p); err != nil {
// 			log.Println(err)
// 			return
// 		}

//		}
//	}
func (h *Handler) uploadVideoHandlerfunc(c *gin.Context) {
	// Parse the form data including the uploaded file
	// form, err := c.MultipartForm()
	// if err != nil {
	// 	c.String(http.StatusBadRequest, fmt.Sprintf("Error: %s", err.Error()))
	// 	return
	// }
	// log.Println("context file", form)

	// Get the file from the form
	_, file, err := c.Request.FormFile("video")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("Error: %s", err.Error()))
		return
	}

	// Save the video to the "uploads" folder
	savePath := "uploads"
	if err := saveUploadedFile(file, savePath, c); err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("Error saving file: %s", err.Error()))
		return
	}
	c.JSON(http.StatusOK, "File uploaded successfully!")
}
func saveUploadedFile(file *multipart.FileHeader, savePath string, c *gin.Context) error {
	// Create the "uploads" folder if it doesn't exist
	if _, err := os.Stat(savePath); os.IsNotExist(err) {
		err := os.Mkdir(savePath, os.ModePerm)
		if err != nil {
			return fmt.Errorf("failed to create folder: %s", err)
		}
	}

	// Generate a unique filename for the uploaded video
	filename := filepath.Join(savePath, file.Filename)

	// Save the file to the server
	if err := c.SaveUploadedFile(file, filename); err != nil {
		return fmt.Errorf("failed to save file: %s", err)
	}

	return nil
}
func (h *Handler) addBranchHandler(c *gin.Context) {
	req := models.AddbranchParams{}

	if err := c.Bind(&req); err != nil {
		err := apperrors.NewExpectationFailed("error binding branch params!!")
		c.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}
	result, err := h.AdminService.AddbranchService(c.Request.Context(), req)
	if err != nil {
		err := apperrors.NewInternal()
		c.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": result,
	})

}

func (h *Handler) deleteBranchHandler(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		err := apperrors.NewExpectationFailed("id is required param!!")
		c.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}

}

func (h *Handler) getBranchHandler(c *gin.Context) {
	req := models.GetBranchParams{}
	if err := c.Bind(&req); err != nil {
		err := apperrors.NewExpectationFailed("branch parametres are required params !!")
		c.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}
}

func (h *Handler) getAllBranchHandler(c *gin.Context) {
	id := c.Param("code")
	if id == "" {
		err := apperrors.NewExpectationFailed("id is required param!!")
		c.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}
	branches, err := h.AdminService.GetAllBranchesService(c.Request.Context(), id)
	if err != nil {
		log.Println("error occured getting branches", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, branches)
}

func (h *Handler) updateBranchHandler(c *gin.Context) {
	req := models.AddbranchParams{}

	if err := c.Bind(&req); err != nil {
		err := apperrors.NewExpectationFailed("error binding branch params!!")
		c.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}
}

func (h *Handler) AddCounterHandler(c *gin.Context) {
	req := models.AddCounterParams{}

	if err := c.Bind(&req); err != nil {
		err := apperrors.NewExpectationFailed("error binding counter params!!")
		c.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}
	result, err := h.AdminService.AddCounterService(c.Request.Context(), req)
	if err != nil {
		err := apperrors.NewInternal()
		c.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": result,
	})

}
func (h *Handler) DeleteCounterHandler(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		err := apperrors.NewExpectationFailed("id is required param!!")
		c.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}
	result, err := h.AdminService.DeleteCounterService(c.Request.Context(), id)
	if err != nil {
		err := apperrors.NewInternal()
		c.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}
	res, err := h.LicenseService.DeleteCounterFromLicense(id)
	if err != nil {
		err := apperrors.NewInternal()
		c.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}
	if res {
		c.JSON(http.StatusOK, gin.H{
			"message": result,
		})
	} else {
		err := apperrors.NewInternal()
		c.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}
}
func (h *Handler) GetAllCountersHandler(c *gin.Context) {
	id := c.Param("code")
	if id == "" {
		err := apperrors.NewExpectationFailed("id is required param!!")
		c.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}
	result, err := h.AdminService.GetAllCountersService(c.Request.Context(), id)
	if err != nil {
		err := apperrors.NewInternal()
		c.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": result,
	})

}
func (h *Handler) GetAllUnActivatedCountersHandler(c *gin.Context) {
	id := c.Param("code")
	if id == "" {
		err := apperrors.NewExpectationFailed("id is required param!!")
		c.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}
	result, err := h.AdminService.GetAllUnActivatedCountersService(c.Request.Context(), id)
	if err != nil {
		err := apperrors.NewInternal()
		c.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": result,
	})

}
func (h *Handler) GetCounterHandler(c *gin.Context) {
	req := models.GetCounterParams{}

	if err := c.Bind(&req); err != nil {
		err := apperrors.NewExpectationFailed("error binding counter params!!")
		c.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}
	result, err := h.AdminService.GetCounterService(c.Request.Context(), req)
	if err != nil {
		err := apperrors.NewInternal()
		c.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": result,
	})

}
func (h *Handler) UpdateCounterHandler(c *gin.Context) {
	req := models.UpdateCounterParams{}

	if err := c.Bind(&req); err != nil {
		err := apperrors.NewExpectationFailed("error binding counter params!!")
		c.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}
	err := h.AdminService.UpdateCounterService(c.Request.Context(), req)
	if err != nil {
		err := apperrors.NewInternal()
		c.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})

}

func (h *Handler) AddSectionHandler(c *gin.Context) {
	req := models.AddSectionParams{}

	if err := c.Bind(&req); err != nil {
		err := apperrors.NewExpectationFailed("error binding section params!!")
		c.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}

}
func (h *Handler) DeleteSectionHandler(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		err := apperrors.NewExpectationFailed("id is required param!!")
		c.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}

}
func (h *Handler) GetAllSectionsService(c *gin.Context) {
	id := c.Param("code")
	if id == "" {
		err := apperrors.NewExpectationFailed("code is required param!!")
		c.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}
}
func (h *Handler) GetSectionHandler(c *gin.Context) {
	req := models.GetSectionParams{}

	if err := c.Bind(&req); err != nil {
		err := apperrors.NewExpectationFailed("error binding section params!!")
		c.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}
}
func (h *Handler) UpdateSectionHandler(c *gin.Context) {
	req := models.UpdateSectionParams{}

	if err := c.Bind(&req); err != nil {
		err := apperrors.NewExpectationFailed("error binding section params!!")
		c.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}
}

func (h *Handler) AddServiceHandler(c *gin.Context) {
	req := models.GetServiceResult{}

	if err := c.Bind(&req); err != nil {
		log.Println("error binding params reason:", err)
		err := apperrors.NewExpectationFailed("error binding service params!!")
		c.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}
	result, err := h.AdminService.AddService(c.Request.Context(), req)
	if err != nil {
		log.Println("error adding service:", err)
		err := apperrors.NewInternal()
		c.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": result,
	})
}
func (h *Handler) DeleteServiceHandler(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		err := apperrors.NewExpectationFailed("id is required param!!")
		c.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}
	res, err := h.AdminService.DeleteService(c.Request.Context(), id)
	if err != nil {
		log.Println("error occured while deleting service")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, res)
}
func (h *Handler) GetAllServicesHandler(c *gin.Context) {
	id := c.Param("code")
	if id == "" {
		err := apperrors.NewExpectationFailed("code is required param!!")
		c.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}

	result, err := h.AdminService.GetAllServices(c.Request.Context(), id)
	if err != nil {
		log.Println("error getting service:", err)
		err := apperrors.NewInternal()
		c.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": result,
	})

}
func (h *Handler) GetServiceHandler(c *gin.Context) {
	req := models.GetServiceParams{}

	if err := c.Bind(&req); err != nil {
		err := apperrors.NewExpectationFailed("error binding service params!!")
		c.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}
	log.Println("params:::::", req)
	result, err := h.AdminService.GetService(c.Request.Context(), req)
	if err != nil {
		log.Println("error getting service:", err)
		err := apperrors.NewInternal()
		c.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": result,
	})
}
func (h *Handler) UpdateServiceHandler(c *gin.Context) {
	req := models.UpdateServiceParams{}

	if err := c.Bind(&req); err != nil {
		log.Println("error binding data", err)
		err := apperrors.NewExpectationFailed("error binding service params!!")
		c.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}
	err := h.AdminService.UpdateService(context.TODO(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, "ok")
}

func (h *Handler) AddUserHandler(c *gin.Context) {
	req := models.AddUserParams{}
	if err := c.Bind(&req); err != nil {
		log.Println(err)
		err := apperrors.NewExpectationFailed("error binding service params!!")
		c.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}
	log.Println("params:::::", req)
	result, err := h.AdminService.AddUserService(c.Request.Context(), req)
	if err != nil {
		log.Println("error adding user:", err)
		err := apperrors.NewInternal()
		c.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": result,
	})

}
func (h *Handler) DeleteUserHandler(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		err := apperrors.NewExpectationFailed("code is required param!!")
		c.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}
	log.Println("params:::::", id)
	result, err := h.AdminService.DeleteUserService(c.Request.Context(), id)
	if err != nil {
		log.Println("error deleting user:", err)
		err := apperrors.NewInternal()
		c.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": result,
	})
}
func (h *Handler) GetAllUsersHandler(c *gin.Context) {
	req := models.GetAllUsersParams{}
	if err := c.Bind(&req); err != nil {
		log.Println(err)
		err := apperrors.NewExpectationFailed("error binding service params!!")
		c.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}
	log.Println("params:::::", req)
	result, err := h.AdminService.GetAllUsers(c.Request.Context(), req)
	if err != nil {
		log.Println("error getting all users:", err)
		err := apperrors.NewInternal()
		c.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": result,
	})

}
func (h *Handler) GetUserHandler(c *gin.Context) {
	req := models.GetUserParams{}
	if err := c.Bind(&req); err != nil {
		log.Println(err)
		err := apperrors.NewExpectationFailed("error binding service params!!")
		c.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}
	log.Println("params:::::", req)
	result, err := h.AdminService.GetUser(c.Request.Context(), req)
	if err != nil {
		log.Println("error getting user:", err)
		err := apperrors.NewInternal()
		c.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": result,
	})

}
func (h *Handler) UpdateUserEmailHandler(c *gin.Context) {
	req := models.UpdateUserEmailParams{}
	if err := c.Bind(&req); err != nil {
		log.Println(err)
		err := apperrors.NewExpectationFailed("error binding service params!!")
		c.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}
	log.Println("params:::::", req)
	err := h.AdminService.UpdateUserEmailService(c.Request.Context(), req)
	if err != nil {
		log.Println("error adding user:", err)
		err := apperrors.NewInternal()
		c.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": true,
	})

}
func (h *Handler) UpdateUserPasswordHandler(c *gin.Context) {
	req := models.UpdateUserPasswordParams{}
	if err := c.Bind(&req); err != nil {
		log.Println(err)
		err := apperrors.NewExpectationFailed("error binding service params!!")
		c.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}
	log.Println("params:::::", req)
	err := h.AdminService.UpdateUserPassword(c.Request.Context(), req)
	if err != nil {
		log.Println("error updating password:", err)
		err := apperrors.NewInternal()
		c.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": true,
	})

}
func (h *Handler) UpdateUserHandler(c *gin.Context) {
	req := models.UpdateUserParams{}
	if err := c.Bind(&req); err != nil {
		log.Println(err)
		err := apperrors.NewExpectationFailed("error binding service params!!")
		c.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}
	log.Println("params:::::", req)
	err := h.AdminService.UpdateUserService(c.Request.Context(), req)
	if err != nil {
		log.Println("error updating user:", err)
		err := apperrors.NewInternal()
		c.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": true,
	})
}
