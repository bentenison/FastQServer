package models

import "github.com/gocraft/dbr/v2"

type Video struct {
	Id          string `json:"id"`
	FileName    string `json:"file_name"`
	Type        string `json:"type"`
	Size        int    `json:"size"`
	ModifiedAt  string `json:"modified_at"`
	CreatedAt   string `json:"created_at"`
	UpdatedBy   string `json:"updated_by"`
	BranchCode  string `json:"branch_code"`
	CompanyCode string `json:"company_code"`
	CompanyName string `json:"company_name"`
	BranchName  string `json:"branch_name"`
	IsDeleted   bool   `json:"isdeleted"`
	Day         string `json:"day"`
	Ordinality  string `json:"ordinality"`
}
type AnnouncementConf struct {
	Id          string `json:"id"`
	Text        string `json:"announce_text"`
	Speed       int    `json:"speed"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	UpdatedBy   string `json:"updated_by"`
	BranchCode  string `json:"branch_code"`
	CompanyCode string `json:"company_code"`
}

// CREATE TABLE `announce_text` (

// 	PRIMARY KEY (`id`)
//   ) ;
type VideoSchedular struct {
	Id          string `json:"id"`
	VideoId     string `json:"video_id"`
	Day         string `json:"day"`
	Ordinality  int    `json:"ordinality"`
	CreatedAt   string `json:"created_at"`
	CreatedBy   string `json:"created_by"`
	UpdatedAt   string `json:"updated_at"`
	UpdatedBy   string `json:"updated_by"`
	BranchCode  string `json:"branch_code"`
	CompanyCode string `json:"company_code"`
}

type AudioConfig struct {
	ID           string `json:"id"`
	HideDate     bool   `json:"hide_date"`
	Message      string `json:"message"`
	TTSLang      string `json:"tts_lang"`
	RepeatCall   int    `json:"repeat_call"`
	BranchCode   string `json:"branch_code"`
	BranchName   string `json:"branch_name"`
	CompanyCode  string `json:"company_code"`
	CompanyId    string `json:"company_id"`
	CompanyName  string `json:"company_name"`
	CreatedAt    string `json:"created_at"`
	UpdateAt     string `json:"updated_at"`
	UpdatedBy    string `json:"updated_by"`
	BellName     string `json:"bell_name"`
	StartingBell bool   `json:"starting_bell"`
	EndingBell   bool   `json:"ending_bell"`
}

type DSConfig struct {
	Id           string `json:"id"`
	BranchCode   string `json:"branch_code"`
	BranchName   string `json:"branch_name"`
	CompnyCode   string `json:"company_code"`
	CompanyName  string `json:"company_name"`
	Template     string `json:"template"`
	Retain_queue bool   `json:"retain_q"`
	ConfirmDelay bool   `json:"confirmation_delay"`
	ShowForm     bool   `json:"show_form"`
	ShowScroll   bool   `json:"show_scroll"`
	ShowDate     bool   `json:"show_dt"`
	Volume       int    `json:"volume"`
	FormConfigId string `json:"form_conf"`
	OnlineForm   string `json:"online_form"`
	ShowURL      bool   `json:"show_url"`
	URL          string `json:"url"`
}
type FormConfig struct {
	Name        bool   `json:"name"`
	Phone       bool   `json:"phone"`
	Email       bool   `json:"email"`
	SendMail    bool   `json:"send_mail"`
	Date        bool   `json:"date"`
	CreatedAt   string `json:"created_at"`
	CreatedBy   string `json:"created_by"`
	UpdatedAt   string `json:"updated_at"`
	UpdatedBy   string `json:"updated_by"`
	CompanyCode string `json:"company_code"`
	CompanyName string `json:"company_name"`
	BranchCode  string `json:"branch_code"`
	BranchName  string `json:"branch_name"`
	ID          string `json:"id"`
}

type TicketConf struct {
	ImageHeader    bool   `json:"image_header" db:"image_header"`
	ImageURL       string `json:"image_url" db:"image_url"`
	NoShowExp      int    `json:"no_show_exp" db:"no_show_exp"`
	TextHeader     bool   `json:"text_header" db:"text_header"`
	HeaderText     string `json:"header_text" db:"header_text"`
	HideTime       bool   `json:"hide_time" db:"hide_time"`
	PrintDuplicate bool   `json:"print_duplicate" db:"print_duplicate"`
	DateFormat     string `json:"date_format" db:"date_format"`
	PrintPosition  bool   `json:"print_position" db:"print_position"`
	NumberDigits   int    `json:"number_digits" db:"number_digits"`
	EstimatedTime  bool   `json:"estimated_time" db:"estimated_time"`
	PopUpDelay     bool   `json:"pop_up_delay" db:"pop_up_delay"`
	CreatedAt      string `json:"created_at" db:"created_at"`
	UpdatedAt      string `json:"updated_at" db:"updated_at"`
	UpdatedBy      string `json:"updated_by" db:"updated_by"`
	ID             string `json:"id" db:"id"`
	BranchCode     string `json:"branch_code" db:"branch_code"`
	BranchName     string `json:"branch_name" db:"branch_name"`
	CompanyCode    string `json:"company_code" db:"company_code"`
	CompanyName    string `json:"company_name" db:"company_name"`
}
type VideoRes struct {
	Id          dbr.NullString `json:"id"`
	FileName    dbr.NullString `json:"file_name"`
	Type        dbr.NullString `json:"type"`
	Size        dbr.NullInt64  `json:"size"`
	ModifiedAt  dbr.NullString `json:"modified_at"`
	CreatedAt   dbr.NullString `json:"created_at"`
	UpdatedBy   dbr.NullString `json:"updated_by"`
	BranchCode  dbr.NullString `json:"branch_code"`
	CompanyCode dbr.NullString `json:"company_code"`
	CompanyName dbr.NullString `json:"company_name"`
	BranchName  dbr.NullString `json:"branch_name"`
	IsDeleted   dbr.NullBool   `json:"isdeleted"`
	Day         dbr.NullString `json:"day"`
	Ordinality  dbr.NullString `json:"ordinality"`
}

type VideoSchedularRes struct {
	Id          dbr.NullString `json:"id"`
	VideoId     dbr.NullString `json:"video_id"`
	Day         dbr.NullString `json:"day"`
	Ordinality  dbr.NullInt64  `json:"ordinality"`
	CreatedAt   dbr.NullString `json:"created_at"`
	CreatedBy   dbr.NullString `json:"created_by"`
	UpdatedAt   dbr.NullString `json:"updated_at"`
	UpdatedBy   dbr.NullString `json:"updated_by"`
	BranchCode  dbr.NullString `json:"branch_code"`
	CompanyCode dbr.NullString `json:"company_code"`
}

type AudioConfigRes struct {
	ID           dbr.NullString `json:"id"`
	HideDate     dbr.NullBool   `json:"hide_date"`
	Message      dbr.NullString `json:"message"`
	TTSLang      dbr.NullString `json:"tts_lang"`
	RepeatCall   dbr.NullInt64  `json:"repeat_call"`
	BranchCode   dbr.NullString `json:"branch_code"`
	BranchName   dbr.NullString `json:"branch_name"`
	CompanyCode  dbr.NullString `json:"company_code"`
	CompanyId    dbr.NullString `json:"company_id"`
	CompanyName  dbr.NullString `json:"company_name"`
	CreatedAt    dbr.NullString `json:"created_at"`
	UpdateAt     dbr.NullString `json:"updated_at"`
	UpdatedBy    dbr.NullString `json:"updated_by"`
	BellName     dbr.NullString `json:"bell_name"`
	StartingBell dbr.NullBool   `json:"starting_bell"`
	EndingBell   dbr.NullBool   `json:"ending_bell"`
}

type DSConfigRes struct {
	Id           dbr.NullString `json:"id"`
	BranchCode   dbr.NullString `json:"branch_code"`
	BranchName   dbr.NullString `json:"branch_name"`
	CompnyCode   dbr.NullString `json:"company_code"`
	CompanyName  dbr.NullString `json:"company_name"`
	Template     dbr.NullString `json:"template"`
	Retain_queue dbr.NullBool   `json:"retain_q"`
	ConfirmDelay dbr.NullBool   `json:"confirmation_delay"`
	ShowForm     dbr.NullBool   `json:"show_form"`
	ShowScroll   dbr.NullBool   `json:"show_scroll"`
	ShowDate     dbr.NullBool   `json:"show_dt"`
	Volume       dbr.NullInt64  `json:"volume"`
	FormConfigId dbr.NullString `json:"form_conf"`
	OnlineForm   dbr.NullString `json:"online_form"`
	ShowURL      dbr.NullBool   `json:"show_url"`
	URL          dbr.NullString `json:"url"`
}
type FormConfigRes struct {
	Name        dbr.NullBool   `json:"name"`
	Phone       dbr.NullBool   `json:"phone"`
	Email       dbr.NullBool   `json:"email"`
	SendMail    dbr.NullBool   `json:"send_mail"`
	Date        dbr.NullBool   `json:"date"`
	CreatedAt   dbr.NullString `json:"created_at"`
	CreatedBy   dbr.NullString `json:"created_by"`
	UpdatedAt   dbr.NullString `json:"updated_at"`
	UpdatedBy   dbr.NullString `json:"updated_by"`
	CompanyCode dbr.NullString `json:"company_code"`
	CompanyName dbr.NullString `json:"company_name"`
	BranchCode  dbr.NullString `json:"branch_code"`
	BranchName  dbr.NullString `json:"branch_name"`
	ID          dbr.NullString `json:"id"`
}
type AnnouncementConfRes struct {
	Id          dbr.NullString `json:"id"`
	Text        dbr.NullString `json:"announce_text"`
	Speed       dbr.NullInt64  `json:"speed"`
	CreatedAt   dbr.NullString `json:"created_at"`
	UpdatedAt   dbr.NullString `json:"updated_at"`
	UpdatedBy   dbr.NullString `json:"updated_by"`
	BranchCode  dbr.NullString `json:"branch_code"`
	CompanyCode dbr.NullString `json:"company_code"`
	IsSelected  dbr.NullBool   `json:"isSelected"`
}
type TicketConfRes struct {
	ImageHeader    dbr.NullBool   `json:"image_header" db:"image_header"`
	ImageURL       dbr.NullString `json:"image_url" db:"image_url"`
	NoShowExp      dbr.NullInt64  `json:"no_show_exp" db:"no_show_exp"`
	TextHeader     dbr.NullBool   `json:"text_header" db:"text_header"`
	HeaderText     dbr.NullString `json:"header_text" db:"header_text"`
	HideTime       dbr.NullBool   `json:"hide_time" db:"hide_time"`
	PrintDuplicate dbr.NullBool   `json:"print_duplicate" db:"print_duplicate"`
	DateFormat     dbr.NullString `json:"date_format" db:"date_format"`
	PrintPosition  dbr.NullBool   `json:"print_position" db:"print_position"`
	NumberDigits   dbr.NullInt64  `json:"number_digits" db:"number_digits"`
	EstimatedTime  dbr.NullBool   `json:"estimated_time" db:"estimated_time"`
	PopUpDelay     dbr.NullBool   `json:"pop_up_delay" db:"pop_up_delay"`
	CreatedAt      dbr.NullString `json:"created_at" db:"created_at"`
	UpdatedAt      dbr.NullString `json:"updated_at" db:"updated_at"`
	UpdatedBy      dbr.NullString `json:"updated_by" db:"updated_by"`
	ID             string         `json:"id" db:"id"`
	BranchCode     dbr.NullString `json:"branch_code" db:"branch_code"`
	BranchName     dbr.NullString `json:"branch_name" db:"branch_name"`
	CompanyCode    dbr.NullString `json:"company_code" db:"company_code"`
	CompanyName    dbr.NullString `json:"company_name" db:"company_name"`
}

type AllConfig struct {
	VideoConf        []VideoRes          `json:"video_conf,omitempty"`
	SchedularConf    []VideoSchedularRes `json:"schedular_conf,omitempty"`
	AudioConf        AudioConfigRes      `json:"audio_conf,omitempty"`
	DSConf           DSConfigRes         `json:"ds_conf,omitempty"`
	AnnouncementConf AnnouncementConfRes `json:"announcement_conf,omitempty"`
	TicketConf       TicketConfRes       `json:"ticket_conf,omitempty"`
}

// CREATE TABLE `announce_text` (
// 	`id`
// 	`announce_text`
// 	`speed` int(255) DEFAULT 0,
// 	`created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
// 	`updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
// 	`created_by`
// 	`updated_by`
// 	`branch_code`
// 	`company_code`
// 	PRIMARY KEY (`id`)
//   ) ;

type CounterService struct {
	ID        int    `json:"CounterServiceID,omitempty"`
	CounterID string `json:"CounterID,omitempty"`
	ServiceID string `json:"ServiceID,omitempty"`
}
