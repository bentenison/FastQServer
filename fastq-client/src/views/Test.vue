<template>
  <div>
    <!-- Your Vue.js component content here -->

    <!-- Trigger print without a button click for demonstration purposes -->
    <v-btn @click="print">Print Document</v-btn>
    <v-btn @click="printElectronTicket">Print Test Document</v-btn>
  </div>
</template>
  
  <script>
import axios from "axios";
export default {
  methods: {
    printDocument() {
      // Create a hidden iframe
      const iframe = document.createElement("iframe");
      iframe.style.display = "none";
      document.body.appendChild(iframe);

      // Set iframe content with your printable document
      iframe.contentDocument.write(
        "<html><head><title>Your Document</title></head><body>"
      );
      // Append the content you want to print
      iframe.contentDocument.write("<h1>Hello, World!</h1>");
      iframe.contentDocument.write("</body></html>");
      //   if (window.matchMedia) {
      //     var mediaQueryList = window.matchMedia("print");
      //     mediaQueryList.addEventListener(function (mql) {
      //       if (!mql.matches) {
      //         // The print preview has been closed, trigger printing
      //         window.print();
      //       }
      //     });
      //   }

      // Open the print preview
      // window.print();
      // Trigger the print operation

      iframe.contentWindow.print();

      // Remove the iframe after printing
      setTimeout(() => {
        document.body.removeChild(iframe);
      }, 1000); // Adjust the delay based on your content and printing speed
    },
    print() {
      const printContent = `
  <html>
    <body>
      <h1>Hello, Printer!</h1>
    </body>
  </html>
`;

      const printerApiEndpoint = "http://localhost:9100/"; // Replace with the actual API endpoint

      axios
        .get(printerApiEndpoint, printContent)
        .then((response) => {
          console.log("Print request sent successfully:", response.data);
        })
        .catch((error) => {
          console.error("Error sending print request:", error.message);
        });
    },
    wsPrint() {
      let data = new TextEncoder("Hello Worrld!!!");
      var socket = new WebSocket("ws://127.0.0.1:9100/");
      socket.bufferType = "arraybuffer";
      socket.onerror = function (error) {
        alert("Error");
      };
      socket.onopen = function () {
        socket.send(data);
        socket.close(1000, "Work complete");
      };
    },
    async test() {
      // Get a list of attached printers
      var printers = await securePOSPrinter.getPrinters();

      // Print some data

      var data = [
        {
          type: "text",
          value: "Testing 123",
        },
        {
          type: "qrCode",
          value: "https://www.saltlakefilmsociety.org/",
          options: {
            scale: 4,
          },
        },
        {
          type: "barCode",
          value: 123456789999,
        },
        {
          // Path to image on file system
          // mimeType is required when loading images from the file system
          type: "image",
          path: "./assets/printer-image.png",
          mimeType: "image/png",
        },
        {
          // You can also specify a "src" attribute directly
          type: "image",
          src: "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAgAAAAICAIAAABLbSncAAAACXBIWXMAAAsTAAALEwEAmpwYAAAAB3RJTUUH5gIREQQbbUH0CgAAAB1pVFh0Q29tbWVudAAAAAAAQ3JlYXRlZCB3aXRoIEdJTVBkLmUHAAAAMUlEQVQI13WOQQ4AIAzCWv//53kwmjiFYyGAAFBVbKmAjR7Ply6NK6Ufo7elqjye7k6T7hILeH2VsAAAAABJRU5ErkJggg==",
          sectionStyle: {
            alignItems: "center",
            justifyContent: "center",
          },
          attributes: {
            width: "100px",
            height: "100px",
          },
        },
        {
          type: "table",
          header: ["Item", "Price"],
          rows: [
            ["Popcorn", "$5.00"],
            ["Soda", "$3.00"],
            ["Candy", "$2.00"],
          ],
          footer: ["Total", "$10.00"],
        },
      ];

      var options = {
        printerName: "SomePrinterName",
        // You can specify your own global stylesheet if you want
        styleSheet: "body,#container { width: '110mm'}",
        preview: true,
      };

      // Note: This promise will not resolve until
      // 1) The preview window is shown, or
      // 2) The document has been sent to the printer, or
      // 3) An error occurred
      await securePOSPrinter.print(data, options);
    },
    printElectronTicket() {
      let ticketPayload = {
        service: "Parts Order",
        ticket_name: "A-001",
        Waiting: 34,
        timeDate: new Date(),
      };
      axios
        .post("http://localhost:4000/api/print", ticketPayload)
        .then((res) => {
          console.log("res:::::", res);
          this.$toast.success("ticket printed successfully.");
          this.branch = {};
        })
        .catch((err) => {
          console.log(err.response);
          this.$toast.error("error occured while printing ticket!!!");
        });
    },
  },
  mounted() {
    console.log("Primter", window);
  },
};
</script>

  <style>
@media print {
  /* Your print-specific styles here */
  body {
    visibility: hidden;
  }
  #printableContent {
    visibility: visible;
  }
}
</style>
  