package services

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/bentenison/fastq-server/api/helpers"
	"github.com/bentenison/fastq-server/models"
	"github.com/google/uuid"
)

type configService struct {
	db *sql.DB
}

/*
* CoConfig will hold repositories that will eventually be injected into this service layer
* now it means the CoConfig is used to hide the implementation of service layer
 */
type CoConfig struct {
	DB *sql.DB
}

func NewConfigService(c *CoConfig) models.ConfigService {
	return &configService{
		db: c.DB,
	}
}

func (c *configService) AddVideoService(ctx context.Context, v models.Video) (sql.Result, error) {
	v.Id = uuid.NewString()
	t := time.Now()
	tm := t.Format("2006-01-02 15:04:05")
	v.CreatedAt = tm
	// v. = tm
	return AddVideo(ctx, c.db, v)
}
func (c *configService) GetAllVideoService(ctx context.Context, id string) ([]models.VideoRes, error) {
	return getAllVideos(ctx, c.db, id)
}
func (c *configService) GetVideoService(ctx context.Context, id string) (models.VideoRes, error) {
	return getVideo(ctx, c.db, id)
}
func (c *configService) AddSchedularConfigService(ctx context.Context, v models.VideoSchedular) (sql.Result, error) {
	v.Id = uuid.NewString()
	t := time.Now()
	tm := t.Format("2006-01-02 15:04:05")
	v.CreatedAt = tm
	v.UpdatedAt = tm
	return AddVideoSchedular(ctx, c.db, v)
}
func (c *configService) GetSchedularConfService(ctx context.Context, id string) (models.VideoSchedularRes, error) {
	return getVideoSchedular(ctx, c.db, id)
}
func (c *configService) GetAllSchedularConfService(ctx context.Context, id string) ([]models.VideoSchedularRes, error) {
	return getAllVideosSchedular(ctx, c.db, id)
}
func (c *configService) AddAudioConfigService(a models.AudioConfig) error {
	a.ID = uuid.NewString()
	t := time.Now()
	tm := t.Format("2006-01-02 15:04:05")
	a.CreatedAt = tm
	a.UpdateAt = tm
	return insertAudioConfig(c.db, a)
}
func (c *configService) UpdateAudioConfigService(a models.AudioConfig) error {

	t := time.Now()
	tm := t.Format("2006-01-02 15:04:05")
	a.UpdateAt = tm
	a.CreatedAt = tm
	return updateAudioConfig(c.db, a)
}
func (c *configService) GetAudionConfigService(id string) (models.AudioConfigRes, error) {
	return selectOneAudioConfig(c.db, id)

}
func (c *configService) GetAllAudioConfigService(id string) ([]models.AudioConfigRes, error) {
	return selectAllAudioConfigs(c.db, id)
}
func (c *configService) AddDSConfigService(a models.DSConfig) error {
	a.Id = uuid.NewString()
	// t := time.Now()
	// tm := t.Format("2006-01-02 15:04:05")
	// a.Cr = tm
	return insertDSConfig(c.db, a)
}
func (c *configService) UpdateDsConfigService(a models.DSConfig) error {
	// a.Id = uuid.NewString()
	// t := time.Now()
	// tm := t.Format("2006-01-02 15:04:05")
	// a.Cr = tm
	return updateDSConfig(c.db, a)
}
func (c *configService) GetDSConfigService(id string) (models.DSConfigRes, error) {
	return selectOneDSConfig(c.db, id)

}
func (c *configService) AddAnnouncementService(a models.AnnouncementConf) error {
	a.Id = uuid.NewString()
	t := time.Now()
	tm := t.Format("2006-01-02 15:04:05")
	a.CreatedAt = tm
	a.UpdatedAt = tm
	return insertAnnouncement(c.db, a)

}
func (c *configService) GetAnnouncement(a string) (models.AnnouncementConfRes, error) {
	return selectOneAnnouncement(c.db, a)

}
func (c *configService) GetAllAnnouncement(a string) ([]models.AnnouncementConfRes, error) {
	return selectAllAnnouncement(c.db, a)

}
func (c *configService) SelectAnnouncementToDisplay(id, code string) error {
	return selectAnnouncement(c.db, id, code)

}
func (c *configService) UpdateAnnouncement(a string, speed int, text string) error {
	return updateAnnouncement(c.db, a, speed, text)
}
func (c *configService) AddTicketConf(tc models.TicketConf) error {
	tc.ID = uuid.NewString()
	t := time.Now()
	tm := t.Format("2006-01-02 15:04:05")
	tc.CreatedAt = tm
	tc.UpdatedAt = tm
	return insertTicketConf(c.db, tc)
}
func (c *configService) GetTicketConf(a string) (models.TicketConfRes, error) {
	return selectOneTicketConfByCompanyCode(c.db, a)
}
func (c *configService) UpdateTicketConf(code string, tc models.TicketConf) error {
	t := time.Now()
	tm := t.Format("2006-01-02 15:04:05")
	tc.UpdatedAt = tm
	return updateAllTicketConfsByCompanyCode(c.db, code, tc)
}
func UpdateServerDetailsService(db *sql.DB, ip string) error {
	// t := time.Now()
	// tm := t.Format("2006-01-02 15:04:05")
	// tc.UpdatedAt = tm
	sd := models.ServerDetails{}
	sd.ServerIP = ip
	cpu, err := helpers.GetCPUId()
	if err != nil {
		log.Println("error while getting cpu id", err)
		return err
	}
	hdd, err := helpers.GetHDDId()
	if err != nil {
		log.Println("error while getting hdd id", err)
		return err
	}
	sd.ServerDiskId = hdd
	sd.ServerCPU = cpu
	return UpdateServerDetails(db, sd)
}
func (c *configService) GetAllConfService(id string) (models.AllConfig, error) {
	// t := time.Now()
	// tm := t.Format("2006-01-02 15:04:05")
	// tc.UpdatedAt = tm

	return GetAllConf(context.TODO(), c.db, id)
}
func (c *configService) GetServerByIDService(id string) (models.ServerDetails, error) {
	// t := time.Now()
	// tm := t.Format("2006-01-02 15:04:05")
	// tc.UpdatedAt = tm

	return GetServerDetailsBycompany(id, c.db)
}

func AddVideo(ctx context.Context, db *sql.DB, newVideo models.Video) (sql.Result, error) {
	query := "INSERT INTO videos (id, file_name, type, size, modified_at, created_at, " +
		"updated_by, branch_code, company_code, company_name, branch_name, isdeleted) " +
		"VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"

	// Execute the SQL query
	res, err := db.Exec(query,
		newVideo.Id,
		newVideo.FileName,
		newVideo.Type,
		newVideo.Size,
		newVideo.ModifiedAt,
		newVideo.CreatedAt,
		newVideo.UpdatedBy,
		newVideo.BranchCode,
		newVideo.CompanyCode,
		newVideo.CompanyName,
		newVideo.BranchName,
		newVideo.IsDeleted,
	)
	if err != nil {
		log.Println(err)
		return res, err
	}
	return res, nil
}

func getVideo(ctx context.Context, db *sql.DB, id string) (models.VideoRes, error) {
	var video models.VideoRes
	query := "SELECT * FROM videos WHERE id = ? LIMIT 1"
	err := db.QueryRow(query, id).Scan(
		&video.Id,
		&video.FileName,
		&video.Type,
		&video.Size,
		&video.ModifiedAt,
		&video.CreatedAt,
		&video.UpdatedBy,
		&video.BranchCode,
		&video.CompanyCode,
		&video.CompanyName,
		&video.BranchName,
		&video.IsDeleted,
	)
	if err == sql.ErrNoRows {
		log.Println("No record found.")
	} else if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Video Record: %+v\n", video)
	}
	return video, nil
}

func getAllVideos(ctx context.Context, db *sql.DB, id string) ([]models.VideoRes, error) {
	var videos []models.VideoRes
	query := "SELECT * FROM videos where company_code =?"
	rows, err := db.Query(query, id)
	if err != nil {
		log.Println(err)
		return videos, err
	}
	defer rows.Close()

	for rows.Next() {
		var video models.VideoRes
		err := rows.Scan(
			&video.Id,
			&video.FileName,
			&video.Type,
			&video.Size,
			&video.ModifiedAt,
			&video.CreatedAt,
			&video.UpdatedBy,
			&video.BranchCode,
			&video.CompanyCode,
			&video.CompanyName,
			&video.BranchName,
			&video.IsDeleted,
			&video.Day,
			&video.Ordinality,
		)
		if err != nil {
			log.Println(err)
			return videos, err
		}
		videos = append(videos, video)
	}

	if err := rows.Err(); err != nil {
		log.Println(err)
		return videos, err
	}

	// log.Println("All Video Records:")
	// for _, v := range videos {
	// 	log.Printf("%+v\n", v)
	// }
	return videos, nil
}

func AddVideoSchedular(ctx context.Context, db *sql.DB, newVideoScheduler models.VideoSchedular) (sql.Result, error) {
	query := "INSERT INTO video_schedular (id, video_id, day, ordinality, created_at, created_by, updated_at, updated_by, branch_code, company_code) " +
		"VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"

	// Execute the SQL query
	res, err := db.Exec(query,
		newVideoScheduler.Id,
		newVideoScheduler.VideoId,
		newVideoScheduler.Day,
		newVideoScheduler.Ordinality,
		newVideoScheduler.CreatedAt,
		newVideoScheduler.CreatedBy,
		newVideoScheduler.UpdatedAt,
		newVideoScheduler.UpdatedBy,
		newVideoScheduler.BranchCode,
		newVideoScheduler.CompanyCode,
	)
	if err != nil {
		log.Fatal(err)
		return res, err
	}

	log.Println("VideoScheduler record inserted successfully.")
	return res, nil
}

func getVideoSchedular(ctx context.Context, db *sql.DB, id string) (models.VideoSchedularRes, error) {
	var videoScheduler models.VideoSchedularRes
	query := "SELECT * FROM video_schedulers WHERE id = ? LIMIT 1"
	err := db.QueryRow(query, "1").Scan(
		&videoScheduler.Id,
		&videoScheduler.VideoId,
		&videoScheduler.Day,
		&videoScheduler.Ordinality,
		&videoScheduler.CreatedAt,
		&videoScheduler.CreatedBy,
		&videoScheduler.UpdatedAt,
		&videoScheduler.UpdatedBy,
		&videoScheduler.BranchCode,
		&videoScheduler.CompanyCode,
	)
	if err == sql.ErrNoRows {
		log.Println("No record found.")
	} else if err != nil {
		log.Fatal(err)
		return videoScheduler, err
	} else {
		// log.Printf("VideoScheduler Record: %+v\n", videoScheduler)
	}
	return videoScheduler, nil
}

func getAllVideosSchedular(ctx context.Context, db *sql.DB, id string) ([]models.VideoSchedularRes, error) {
	// Select all records from the video_schedulers table
	var videoSchedulers []models.VideoSchedularRes
	query := "SELECT * FROM video_schedular where company_code =?"
	rows, err := db.Query(query, id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var videoScheduler models.VideoSchedularRes
		err := rows.Scan(
			&videoScheduler.Id,
			&videoScheduler.VideoId,
			&videoScheduler.Day,
			&videoScheduler.Ordinality,
			&videoScheduler.CreatedAt,
			&videoScheduler.CreatedBy,
			&videoScheduler.UpdatedAt,
			&videoScheduler.UpdatedBy,
			&videoScheduler.BranchCode,
			&videoScheduler.CompanyCode,
		)
		if err != nil {
			log.Fatal(err)
		}
		videoSchedulers = append(videoSchedulers, videoScheduler)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	// log.Println("All VideoScheduler Records:")
	// for _, vs := range videoSchedulers {
	// 	log.Printf("%+v\n", vs)
	// }
	return videoSchedulers, nil
}

func insertAudioConfig(db *sql.DB, audioConfig models.AudioConfig) error {
	query := "INSERT INTO audio_conf (id, hide_date, message, tts_lang, repeat_call, " +
		"branch_code, branch_name, company_code, company_id, company_name, created_at, " +
		"updated_at, updated_by, bell_name, starting_bell, ending_bell) " +
		"VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"

	_, err := db.Exec(query,
		audioConfig.ID,
		audioConfig.HideDate,
		audioConfig.Message,
		audioConfig.TTSLang,
		audioConfig.RepeatCall,
		audioConfig.BranchCode,
		audioConfig.BranchName,
		audioConfig.CompanyCode,
		audioConfig.CompanyId,
		audioConfig.CompanyName,
		audioConfig.CreatedAt,
		audioConfig.UpdateAt,
		audioConfig.UpdatedBy,
		audioConfig.BellName,
		audioConfig.StartingBell,
		audioConfig.EndingBell,
	)

	return err
}
func updateAudioConfig(db *sql.DB, audioConfig models.AudioConfig) error {
	query := "update audio_conf set hide_date=?, message=?, tts_lang=?, repeat_call=?, " +
		"branch_code=?, branch_name=?, company_code=?, company_id=?, company_name=?, created_at=?, " +
		"updated_at=?, updated_by=?, bell_name=?, starting_bell=?, ending_bell=? where id = ?"

	_, err := db.Exec(query,
		audioConfig.HideDate,
		audioConfig.Message,
		audioConfig.TTSLang,
		audioConfig.RepeatCall,
		audioConfig.BranchCode,
		audioConfig.BranchName,
		audioConfig.CompanyCode,
		audioConfig.CompanyId,
		audioConfig.CompanyName,
		audioConfig.CreatedAt,
		audioConfig.UpdateAt,
		audioConfig.UpdatedBy,
		audioConfig.BellName,
		audioConfig.StartingBell,
		audioConfig.EndingBell,
		audioConfig.ID,
	)

	return err
}

func selectOneAudioConfig(db *sql.DB, id string) (models.AudioConfigRes, error) {
	var audioConfig models.AudioConfigRes
	query := "SELECT * FROM audio_conf WHERE company_code = ? LIMIT 1"

	err := db.QueryRow(query, id).Scan(
		&audioConfig.ID,
		&audioConfig.HideDate,
		&audioConfig.Message,
		&audioConfig.TTSLang,
		&audioConfig.RepeatCall,
		&audioConfig.BranchCode,
		&audioConfig.BranchName,
		&audioConfig.CompanyCode,
		&audioConfig.CompanyId,
		&audioConfig.CompanyName,
		&audioConfig.CreatedAt,
		&audioConfig.UpdateAt,
		&audioConfig.UpdatedBy,
		&audioConfig.BellName,
		&audioConfig.StartingBell,
		&audioConfig.EndingBell,
	)

	return audioConfig, err
}

func selectAllAudioConfigs(db *sql.DB, id string) ([]models.AudioConfigRes, error) {
	var audioConfigs []models.AudioConfigRes
	query := "SELECT * FROM audio_conf where company_code =?"

	rows, err := db.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var audioConfig models.AudioConfigRes
		err := rows.Scan(
			&audioConfig.ID,
			&audioConfig.HideDate,
			&audioConfig.Message,
			&audioConfig.TTSLang,
			&audioConfig.RepeatCall,
			&audioConfig.BranchCode,
			&audioConfig.BranchName,
			&audioConfig.CompanyCode,
			&audioConfig.CompanyId,
			&audioConfig.CompanyName,
			&audioConfig.CreatedAt,
			&audioConfig.UpdateAt,
			&audioConfig.UpdatedBy,
			&audioConfig.BellName,
			&audioConfig.StartingBell,
			&audioConfig.EndingBell,
		)
		if err != nil {
			return nil, err
		}
		audioConfigs = append(audioConfigs, audioConfig)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return audioConfigs, nil
}

// func updateAudioConfig(db *sql.DB, id, newMessage string) error {
// 	query := "UPDATE audio_conf SET message = ?, updated_at = ? WHERE id = ?"

// 	_, err := db.Exec(query, newMessage, time.Now().Format("2006-01-02 15:04:05"), id)
// 	return err
// }

func deleteAudioConfig(db *sql.DB, id string) error {
	query := "DELETE FROM audio_conf WHERE id = ?"

	_, err := db.Exec(query, id)
	return err
}
func updateDSConfig(db *sql.DB, dsConfig models.DSConfig) error {
	query := "update display_conf set branch_code=?, branch_name=?, company_code=?, company_name=?, " +
		"template=?, retain_q=?, confirmation_delay=?, show_form=?, show_scroll=?, show_dt=?, " +
		"volume=?, form_conf=?, online_form=?,show_url=?,url=? where id =?"

	_, err := db.Exec(query,
		dsConfig.BranchCode,
		dsConfig.BranchName,
		dsConfig.CompnyCode,
		dsConfig.CompanyName,
		dsConfig.Template,
		dsConfig.Retain_queue,
		dsConfig.ConfirmDelay,
		dsConfig.ShowForm,
		dsConfig.ShowScroll,
		dsConfig.ShowDate,
		dsConfig.Volume,
		dsConfig.FormConfigId,
		dsConfig.OnlineForm,
		dsConfig.ShowURL,
		dsConfig.URL,
		dsConfig.Id,
	)

	return err
}
func insertDSConfig(db *sql.DB, dsConfig models.DSConfig) error {
	query := "INSERT INTO display_conf (id, branch_code, branch_name, company_code, company_name, " +
		"template, retain_q, confirmation_delay, show_form, show_scroll, show_dt, " +
		"volume, form_conf, online_form,show_url,url) " +
		"VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?,?,?)"

	_, err := db.Exec(query,
		dsConfig.Id,
		dsConfig.BranchCode,
		dsConfig.BranchName,
		dsConfig.CompnyCode,
		dsConfig.CompanyName,
		dsConfig.Template,
		dsConfig.Retain_queue,
		dsConfig.ConfirmDelay,
		dsConfig.ShowForm,
		dsConfig.ShowScroll,
		dsConfig.ShowDate,
		dsConfig.Volume,
		dsConfig.FormConfigId,
		dsConfig.OnlineForm,
		dsConfig.ShowURL,
		dsConfig.URL,
	)

	return err
}

func selectOneDSConfig(db *sql.DB, id string) (models.DSConfigRes, error) {
	var dsConfig models.DSConfigRes
	query := "SELECT * FROM display_conf WHERE company_code = ? LIMIT 1"

	err := db.QueryRow(query, id).Scan(
		&dsConfig.Id,
		&dsConfig.BranchCode,
		&dsConfig.BranchName,
		&dsConfig.CompnyCode,
		&dsConfig.CompanyName,
		&dsConfig.Template,
		&dsConfig.Retain_queue,
		&dsConfig.ConfirmDelay,
		&dsConfig.ShowForm,
		&dsConfig.ShowScroll,
		&dsConfig.ShowDate,
		&dsConfig.Volume,
		&dsConfig.FormConfigId,
		&dsConfig.OnlineForm,
		&dsConfig.ShowURL,
		&dsConfig.URL,
	)

	return dsConfig, err
}

// func updateDSConfig(db *sql.DB, id, newTemplateName string) error {
// 	query := "UPDATE display_conf SET template = ? WHERE company_code = ?"

// 	_, err := db.Exec(query, newTemplateName, id)
// 	return err
// }

func insertFormConfig(db *sql.DB, formConfig models.FormConfig) error {
	query := "INSERT INTO form_configs (name, phone, email, send_mail, date, " +
		"created_at, created_by, updated_at, updated_by, company_code, " +
		"company_name, branch_code, branch_name, id) " +
		"VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"

	_, err := db.Exec(query,
		formConfig.Name,
		formConfig.Phone,
		formConfig.Email,
		formConfig.SendMail,
		formConfig.Date,
		formConfig.CreatedAt,
		formConfig.CreatedBy,
		formConfig.UpdatedAt,
		formConfig.UpdatedBy,
		formConfig.CompanyCode,
		formConfig.CompanyName,
		formConfig.BranchCode,
		formConfig.BranchName,
		formConfig.ID,
	)

	return err
}

func selectOneFormConfig(db *sql.DB, id string) (models.FormConfigRes, error) {
	var formConfig models.FormConfigRes
	query := "SELECT * FROM form_configs WHERE id = ? LIMIT 1"

	err := db.QueryRow(query, id).Scan(
		&formConfig.Name,
		&formConfig.Phone,
		&formConfig.Email,
		&formConfig.SendMail,
		&formConfig.Date,
		&formConfig.CreatedAt,
		&formConfig.CreatedBy,
		&formConfig.UpdatedAt,
		&formConfig.UpdatedBy,
		&formConfig.CompanyCode,
		&formConfig.CompanyName,
		&formConfig.BranchCode,
		&formConfig.BranchName,
		&formConfig.ID,
	)

	return formConfig, err
}

func updateFormConfig(db *sql.DB, id string, newNameSetting bool) error {
	query := "UPDATE form_configs SET name = ? WHERE id = ?"

	_, err := db.Exec(query, newNameSetting, id)
	return err
}

func deleteFormConfig(db *sql.DB, id string) error {
	query := "DELETE FROM form_configs WHERE id = ?"

	_, err := db.Exec(query, id)
	return err
}

func insertAnnouncement(db *sql.DB, announcement models.AnnouncementConf) error {
	query := "INSERT INTO announce_text (id, announce_text, speed, created_at, updated_at, updated_by, branch_code, company_code) " +
		"VALUES (?, ?, ?, ?, ?, ?, ?, ?)"

	_, err := db.Exec(query,
		announcement.Id,
		announcement.Text,
		announcement.Speed,
		announcement.CreatedAt,
		announcement.UpdatedAt,
		announcement.UpdatedBy,
		announcement.BranchCode,
		announcement.CompanyCode,
	)

	return err
}

func selectOneAnnouncement(db *sql.DB, id string) (models.AnnouncementConfRes, error) {
	var announcement models.AnnouncementConfRes
	query := "SELECT * FROM announce_text WHERE company_code = ? and isSelected = 1 LIMIT 1"

	err := db.QueryRow(query, id).Scan(
		&announcement.Id,
		&announcement.Text,
		&announcement.Speed,
		&announcement.CreatedAt,
		&announcement.UpdatedAt,
		&announcement.UpdatedBy,
		&announcement.BranchCode,
		&announcement.CompanyCode,
		&announcement.IsSelected,
	)

	return announcement, err
}
func selectAllAnnouncement(db *sql.DB, id string) ([]models.AnnouncementConfRes, error) {
	var announcements []models.AnnouncementConfRes
	query := "SELECT * FROM announce_text WHERE company_code = ?"
	rows, err := db.Query(query, id)
	if err != nil {
		log.Println("error querying for all anouncements")
		return announcements, err
	}

	defer rows.Close()

	for rows.Next() {
		var announcement models.AnnouncementConfRes
		err := rows.Scan(
			&announcement.Id,
			&announcement.Text,
			&announcement.Speed,
			&announcement.CreatedAt,
			&announcement.UpdatedAt,
			&announcement.UpdatedBy,
			&announcement.BranchCode,
			&announcement.CompanyCode,
			&announcement.IsSelected,
		)
		if err != nil {
			return nil, err
		}
		announcements = append(announcements, announcement)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return announcements, err
}

func insertTicketConf(db *sql.DB, ticketConf models.TicketConf) error {
	_, err := db.Exec(`
		INSERT INTO ticket_conf
		(image_header, image_url, no_show_exp, text_header, header_text, hide_time, print_duplicate,
		date_format, print_position, number_digits, estimated_time, pop_up_delay, created_at,
		updated_at, updated_by, id, branch_code, branch_name, company_code, company_name)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`, ticketConf.ImageHeader, ticketConf.ImageURL, ticketConf.NoShowExp, ticketConf.TextHeader,
		ticketConf.HeaderText, ticketConf.HideTime, ticketConf.PrintDuplicate, ticketConf.DateFormat,
		ticketConf.PrintPosition, ticketConf.NumberDigits, ticketConf.EstimatedTime, ticketConf.PopUpDelay,
		ticketConf.CreatedAt, ticketConf.UpdatedAt, ticketConf.UpdatedBy, ticketConf.ID,
		ticketConf.BranchCode, ticketConf.BranchName, ticketConf.CompanyCode, ticketConf.CompanyName,
	)

	return err
}

func selectOneTicketConfByCompanyCode(db *sql.DB, companyCode string) (models.TicketConfRes, error) {
	var ticketConf models.TicketConfRes
	err := db.QueryRow(`
		SELECT * FROM ticket_conf WHERE company_code = ? LIMIT 1
	`, companyCode).Scan(
		&ticketConf.ImageHeader, &ticketConf.ImageURL, &ticketConf.NoShowExp, &ticketConf.TextHeader,
		&ticketConf.HeaderText, &ticketConf.HideTime, &ticketConf.PrintDuplicate, &ticketConf.DateFormat,
		&ticketConf.PrintPosition, &ticketConf.NumberDigits, &ticketConf.EstimatedTime,
		&ticketConf.PopUpDelay, &ticketConf.CreatedAt, &ticketConf.UpdatedAt, &ticketConf.UpdatedBy,
		&ticketConf.ID, &ticketConf.BranchCode, &ticketConf.BranchName, &ticketConf.CompanyCode,
		&ticketConf.CompanyName,
	)

	return ticketConf, err
}

func selectAllTicketConfsByCompanyCode(db *sql.DB, companyCode string) ([]models.TicketConfRes, error) {
	var ticketConfs []models.TicketConfRes
	rows, err := db.Query(`
		SELECT * FROM ticket_conf WHERE company_code = ?
	`, companyCode)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var ticketConf models.TicketConfRes
		err := rows.Scan(
			&ticketConf.ImageHeader, &ticketConf.ImageURL, &ticketConf.NoShowExp, &ticketConf.TextHeader,
			&ticketConf.HeaderText, &ticketConf.HideTime, &ticketConf.PrintDuplicate, &ticketConf.DateFormat,
			&ticketConf.PrintPosition, &ticketConf.NumberDigits, &ticketConf.EstimatedTime,
			&ticketConf.PopUpDelay, &ticketConf.CreatedAt, &ticketConf.UpdatedAt, &ticketConf.UpdatedBy,
			&ticketConf.ID, &ticketConf.BranchCode, &ticketConf.BranchName, &ticketConf.CompanyCode,
			&ticketConf.CompanyName,
		)
		if err != nil {
			return nil, err
		}
		ticketConfs = append(ticketConfs, ticketConf)
	}

	return ticketConfs, nil
}

func updateAllTicketConfsByCompanyCode(db *sql.DB, companyCode string, updatedTicketConf models.TicketConf) error {
	_, err := db.Exec(`
        UPDATE ticket_conf
        SET image_header = ?, image_url = ?, no_show_exp = ?, text_header = ?, 
            header_text = ?, hide_time = ?, print_duplicate = ?, date_format = ?, 
            print_position = ?, number_digits = ?, estimated_time = ?, pop_up_delay = ?, 
            updated_at = ?, updated_by = ?, id = ?, branch_code = ?, 
            branch_name = ?, company_code = ?, company_name = ?
        WHERE company_code = ?
    `, updatedTicketConf.ImageHeader, updatedTicketConf.ImageURL, updatedTicketConf.NoShowExp,
		updatedTicketConf.TextHeader, updatedTicketConf.HeaderText, updatedTicketConf.HideTime,
		updatedTicketConf.PrintDuplicate, updatedTicketConf.DateFormat, updatedTicketConf.PrintPosition,
		updatedTicketConf.NumberDigits, updatedTicketConf.EstimatedTime, updatedTicketConf.PopUpDelay,
		updatedTicketConf.UpdatedAt, updatedTicketConf.UpdatedBy,
		updatedTicketConf.ID, updatedTicketConf.BranchCode, updatedTicketConf.BranchName,
		updatedTicketConf.CompanyCode, updatedTicketConf.CompanyName, companyCode,
	)

	return err
}

func updateAnnouncement(db *sql.DB, id string, newSpeed int, newText string) error {
	query := "UPDATE announce_text SET speed = ?, announce_text=? WHERE company_code = ?"

	_, err := db.Exec(query, newSpeed, id)
	return err
}
func selectAnnouncement(db *sql.DB, id, code string) error {
	opt := sql.TxOptions{}
	tx, err := db.BeginTx(context.TODO(), &opt)
	if err != nil {
		log.Println("error occured while starting transaction", err)
		return err
	}

	query := "UPDATE announce_text SET isSelected = 1 WHERE company_code = ? and id = ?"
	_, err = tx.Exec("UPDATE announce_text SET isSelected = 0 WHERE isSelected =1")
	_, err = tx.Exec(query, code, id)
	err = tx.Commit()
	if err != nil {
		log.Println("error commiting transaction", err)
		return err
	}
	return err
}

func deleteAnnouncement(db *sql.DB, id string) error {
	query := "DELETE FROM announce_text WHERE company_code = ?"

	_, err := db.Exec(query, id)
	return err
}

func UpdateServerDetails(db *sql.DB, sd models.ServerDetails) error {
	query := "UPDATE server_details SET server_ip = ? WHERE server_cpu = ? and server_disk_id = ?"

	_, err := db.Exec(query, sd.ServerIP, sd.ServerCPU, sd.ServerDiskId)
	return err
}

func GetAllConf(ctx context.Context, db *sql.DB, id string) (models.AllConfig, error) {
	var allconf models.AllConfig
	videos, err := getAllVideos(ctx, db, id)
	if err != nil {
		log.Println("error in getting video conf", err)
		return allconf, err
	}
	allconf.VideoConf = videos
	audioconf, err := selectOneAudioConfig(db, id)
	if err != nil {
		log.Println("error in getting audio conf", err)
		return allconf, err
	}
	allconf.AudioConf = audioconf
	dsconf, err := selectOneDSConfig(db, id)
	if err != nil {
		log.Println("error in getting DS conf", err)
		return allconf, err
	}
	allconf.DSConf = dsconf
	announcement, err := selectOneAnnouncement(db, id)
	if err != nil {
		log.Println("error in getting announcement conf", err)
		return allconf, err
	}
	allconf.AnnouncementConf = announcement
	ticketconf, err := selectOneTicketConfByCompanyCode(db, id)
	if err != nil {
		log.Println("error in getting ticket conf", err)
		// return allconf, err
	}
	allconf.TicketConf = ticketconf
	schedular, err := getAllVideosSchedular(context.TODO(), db, id)
	if err != nil {
		log.Println("error in getting schedular conf", err)
		return allconf, err
	}
	allconf.SchedularConf = schedular
	return allconf, nil
}
func GetServerDetailsBycompany(ID string, db *sql.DB) (models.ServerDetails, error) {

	// Prepare the query
	query := "SELECT server_ip, server_cpu, server_disk_id, id FROM server_details WHERE id = ?"

	// Execute the query
	row := db.QueryRow(query, ID)

	// Scan the result into a ServerDetails struct
	var serverDetails models.ServerDetails
	err := row.Scan(&serverDetails.ServerIP, &serverDetails.ServerCPU, &serverDetails.ServerDiskId, &serverDetails.Id)
	if err != nil {
		if err == sql.ErrNoRows {
			// No rows were returned
			return models.ServerDetails{}, fmt.Errorf("no server details found for server IP: %s", ID)
		}
		return models.ServerDetails{}, err
	}
	return serverDetails, nil
}
