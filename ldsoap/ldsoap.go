package ldsoap

import (
	"encoding/xml"
	"reflect"

	"github.com/fiorix/wsdl2go/soap"
)

// Namespace was auto-generated from WSDL.
var Namespace = "http://Mobile.ws.domoboxbusiness.com/"

// NewMobile creates an initializes a Mobile.
func NewMobile(cli *soap.Client) Mobile {
	return &mobile{cli}
}

// Mobile was auto-generated from WSDL
// and defines interface for the remote service. Useful for testing.
type Mobile interface {
	// RegisterNotifPhoneAndroid was auto-generated from WSDL.
	RegisterNotifPhoneAndroid(α *RegisterNotifPhoneAndroid) (β *RegisterNotifPhoneAndroidResponse, err error)

	// RegisterNotifIpad2 was auto-generated from WSDL.
	RegisterNotifIpad2(α *RegisterNotifIpad2) (β *RegisterNotifIpad2Response, err error)

	// SonosGetAllGroups was auto-generated from WSDL.
	SonosGetAllGroups(α *SonosGetAllGroups) (β *SonosGetAllGroupsResponse, err error)

	// GetVarunaAlarms was auto-generated from WSDL.
	GetVarunaAlarms(α *GetVarunaAlarms) (β *GetVarunaAlarmsResponse, err error)

	// SonosMakeNewGroup was auto-generated from WSDL.
	SonosMakeNewGroup(α *SonosMakeNewGroup) (β *SonosMakeNewGroupResponse, err error)

	// GetDeviceInstantane was auto-generated from WSDL.
	GetDeviceInstantane(α *GetDeviceInstantane) (β *GetDeviceInstantaneResponse, err error)

	// LoginBookmarks was auto-generated from WSDL.
	LoginBookmarks(α *LoginBookmarks) (β *LoginBookmarksResponse, err error)

	// LoginConsumptions was auto-generated from WSDL.
	LoginConsumptions(α *LoginConsumptions) (β *LoginConsumptionsResponse, err error)

	// LoginAlertLogs was auto-generated from WSDL.
	LoginAlertLogs(α *LoginAlertLogs) (β *LoginAlertLogsResponse, err error)

	// GetInstantTypeList was auto-generated from WSDL.
	GetInstantTypeList(α *GetInstantTypeList) (β *GetInstantTypeListResponse, err error)

	// KeepAlive was auto-generated from WSDL.
	KeepAlive(α *KeepAlive) (β *KeepAliveResponse, err error)

	// GetVarunaEvents was auto-generated from WSDL.
	GetVarunaEvents(α *GetVarunaEvents) (β *GetVarunaEventsResponse, err error)

	// GetGroup was auto-generated from WSDL.
	GetGroup(α *GetGroup) (β *GetGroupResponse, err error)

	// SonosSetRandom was auto-generated from WSDL.
	SonosSetRandom(α *SonosSetRandom) (β *SonosSetRandomResponse, err error)

	// SonosDeleteFavorite was auto-generated from WSDL.
	SonosDeleteFavorite(α *SonosDeleteFavorite) (β *SonosDeleteFavoriteResponse, err error)

	// GetDeviceTranslations was auto-generated from WSDL.
	GetDeviceTranslations(α *GetDeviceTranslations) (β *GetDeviceTranslationsResponse, err error)

	// ResetInstantValue was auto-generated from WSDL.
	ResetInstantValue(α *ResetInstantValue) (β *ResetInstantValueResponse, err error)

	// GetDevices was auto-generated from WSDL.
	GetDevices(α *GetDevices) (β *GetDevicesResponse, err error)

	// RegisterNotifTabletteAndroid was auto-generated from WSDL.
	RegisterNotifTabletteAndroid(α *RegisterNotifTabletteAndroid) (β *RegisterNotifTabletteAndroidResponse, err error)

	// GetCameras was auto-generated from WSDL.
	GetCameras(α *GetCameras) (β *GetCamerasResponse, err error)

	// SonosAddAndPlay was auto-generated from WSDL.
	SonosAddAndPlay(α *SonosAddAndPlay) (β *SonosAddAndPlayResponse, err error)

	// GetAlertLogs was auto-generated from WSDL.
	GetAlertLogs(α *GetAlertLogs) (β *GetAlertLogsResponse, err error)

	// GetAlarmFaultsStatistics was auto-generated from WSDL.
	GetAlarmFaultsStatistics(α *GetAlarmFaultsStatistics) (β *GetAlarmFaultsStatisticsResponse, err error)

	// ExecuteAction was auto-generated from WSDL.
	ExecuteAction(α *ExecuteAction) (β *ExecuteActionResponse, err error)

	// Login was auto-generated from WSDL.
	Login(α *Login) (β *LoginResponse, err error)

	// SonosRemoveTrackFromPlaylist was auto-generated from WSDL.
	SonosRemoveTrackFromPlaylist(α *SonosRemoveTrackFromPlaylist) (β *SonosRemoveTrackFromPlaylistResponse, err error)

	// LoginDevicesFromCatg was auto-generated from WSDL.
	LoginDevicesFromCatg(α *LoginDevicesFromCatg) (β *LoginDevicesFromCatgResponse, err error)

	// ClearAndroidNotificationDataBase was auto-generated from WSDL.
	ClearAndroidNotificationDataBase(α *ClearAndroidNotificationDataBase) (β *ClearAndroidNotificationDataBaseResponse, err error)

	// SonosExplore was auto-generated from WSDL.
	SonosExplore(α *SonosExplore) (β *SonosExploreResponse, err error)

	// LoginDevices was auto-generated from WSDL.
	LoginDevices(α *LoginDevices) (β *LoginDevicesResponse, err error)

	// GetHistValue was auto-generated from WSDL.
	GetHistValue(α *GetHistValue) (β *GetHistValueResponse, err error)

	// GetTranslations was auto-generated from WSDL.
	GetTranslations(α *GetTranslations) (β *GetTranslationsResponse, err error)

	// GetStateTranslations was auto-generated from WSDL.
	GetStateTranslations(α *GetStateTranslations) (β *GetStateTranslationsResponse, err error)

	// SonosSetCurrentTrackNumber was auto-generated from WSDL.
	SonosSetCurrentTrackNumber(α *SonosSetCurrentTrackNumber) (β *SonosSetCurrentTrackNumberResponse, err error)

	// SonosDeleteList was auto-generated from WSDL.
	SonosDeleteList(α *SonosDeleteList) (β *SonosDeleteListResponse, err error)

	// GetScenarios was auto-generated from WSDL.
	GetScenarios(α *GetScenarios) (β *GetScenariosResponse, err error)

	// LoginExecuteAction was auto-generated from WSDL.
	LoginExecuteAction(α *LoginExecuteAction) (β *LoginExecuteActionResponse, err error)

	// GetDeviceState was auto-generated from WSDL.
	GetDeviceState(α *GetDeviceState) (β *GetDeviceStateResponse, err error)

	// GetAlarmeAccessLevel was auto-generated from WSDL.
	GetAlarmeAccessLevel(α *GetAlarmeAccessLevel) (β *GetAlarmeAccessLevelResponse, err error)

	// GetAlarmHistoryObjectList was auto-generated from WSDL.
	GetAlarmHistoryObjectList(α *GetAlarmHistoryObjectList) (β *GetAlarmHistoryObjectListResponse, err error)

	// ResetNotif was auto-generated from WSDL.
	ResetNotif(α *ResetNotif) (β *ResetNotifResponse, err error)

	// LoginGroup was auto-generated from WSDL.
	LoginGroup(α *LoginGroup) (β *LoginGroupResponse, err error)

	// GetRTPageDescriptor was auto-generated from WSDL.
	GetRTPageDescriptor(α *GetRTPageDescriptor) (β *GetRTPageDescriptorResponse, err error)

	// SonosGetPlaylist was auto-generated from WSDL.
	SonosGetPlaylist(α *SonosGetPlaylist) (β *SonosGetPlaylistResponse, err error)

	// GetActionTranslations was auto-generated from WSDL.
	GetActionTranslations(α *GetActionTranslations) (β *GetActionTranslationsResponse, err error)

	// RefreshVarunaEvents was auto-generated from WSDL.
	RefreshVarunaEvents(α *RefreshVarunaEvents) (β *RefreshVarunaEventsResponse, err error)

	// RegisterNotif was auto-generated from WSDL.
	RegisterNotif(α *RegisterNotif) (β *RegisterNotifResponse, err error)

	// GetRooms was auto-generated from WSDL.
	GetRooms(α *GetRooms) (β *GetRoomsResponse, err error)

	// RegisterNotifIpad was auto-generated from WSDL.
	RegisterNotifIpad(α *RegisterNotifIpad) (β *RegisterNotifIpadResponse, err error)

	// SonosListenEntryStream was auto-generated from WSDL.
	SonosListenEntryStream(α *SonosListenEntryStream) (β *SonosListenEntryStreamResponse, err error)

	// GetVoiceObjects was auto-generated from WSDL.
	GetVoiceObjects(α *GetVoiceObjects) (β *GetVoiceObjectsResponse, err error)

	// SonosSetRepeat was auto-generated from WSDL.
	SonosSetRepeat(α *SonosSetRepeat) (β *SonosSetRepeatResponse, err error)

	// LoginDeviceState was auto-generated from WSDL.
	LoginDeviceState(α *LoginDeviceState) (β *LoginDeviceStateResponse, err error)

	// LoginInfos was auto-generated from WSDL.
	LoginInfos(α *LoginInfos) (β *LoginInfosResponse, err error)

	// GetPieChartValue was auto-generated from WSDL.
	GetPieChartValue(α *GetPieChartValue) (β *GetPieChartValueResponse, err error)

	// GetCategories was auto-generated from WSDL.
	GetCategories(α *GetCategories) (β *GetCategoriesResponse, err error)

	// CancelMobilePicture was auto-generated from WSDL.
	CancelMobilePicture(α *CancelMobilePicture) (β *CancelMobilePictureResponse, err error)

	// LoginCameras was auto-generated from WSDL.
	LoginCameras(α *LoginCameras) (β *LoginCamerasResponse, err error)

	// SonosAddFavorite was auto-generated from WSDL.
	SonosAddFavorite(α *SonosAddFavorite) (β *SonosAddFavoriteResponse, err error)

	// GetGroups was auto-generated from WSDL.
	GetGroups(α *GetGroups) (β *GetGroupsResponse, err error)

	// LoginScenarios was auto-generated from WSDL.
	LoginScenarios(α *LoginScenarios) (β *LoginScenariosResponse, err error)

	// SonosReorderTrack was auto-generated from WSDL.
	SonosReorderTrack(α *SonosReorderTrack) (β *SonosReorderTrackResponse, err error)

	// GetRTDataTypeList was auto-generated from WSDL.
	GetRTDataTypeList(α *GetRTDataTypeList) (β *GetRTDataTypeListResponse, err error)

	// ApplyMobilePicture was auto-generated from WSDL.
	ApplyMobilePicture(α *ApplyMobilePicture) (β *ApplyMobilePictureResponse, err error)

	// SonosGetLineInList was auto-generated from WSDL.
	SonosGetLineInList(α *SonosGetLineInList) (β *SonosGetLineInListResponse, err error)

	// LoginGroups was auto-generated from WSDL.
	LoginGroups(α *LoginGroups) (β *LoginGroupsResponse, err error)

	// GetTotalDataValue was auto-generated from WSDL.
	GetTotalDataValue(α *GetTotalDataValue) (β *GetTotalDataValueResponse, err error)

	// GetPicture was auto-generated from WSDL.
	GetPicture(α *GetPicture) (β *GetPictureResponse, err error)

	// GetConsumptions was auto-generated from WSDL.
	GetConsumptions(α *GetConsumptions) (β *GetConsumptionsResponse, err error)

	// NewLoginInfos was auto-generated from WSDL.
	NewLoginInfos(α *NewLoginInfos) (β *NewLoginInfosResponse, err error)

	// GetUsers was auto-generated from WSDL.
	GetUsers(α *GetUsers) (β *GetUsersResponse, err error)

	// GetTranslationsREST was auto-generated from WSDL.
	GetTranslationsREST(α *GetTranslationsREST) (β *GetTranslationsRESTResponse, err error)

	// GetBookmarks was auto-generated from WSDL.
	GetBookmarks(α *GetBookmarks) (β *GetBookmarksResponse, err error)

	// GetRoomsNew was auto-generated from WSDL.
	GetRoomsNew(α *GetRoomsNew) (β *GetRoomsNewResponse, err error)

	// GetCategoriesTranslations was auto-generated from WSDL.
	GetCategoriesTranslations(α *GetCategoriesTranslations) (β *GetCategoriesTranslationsResponse, err error)

	// GetDevicesFromCatg was auto-generated from WSDL.
	GetDevicesFromCatg(α *GetDevicesFromCatg) (β *GetDevicesFromCatgResponse, err error)

	// GetRTDataCategoriesList was auto-generated from WSDL.
	GetRTDataCategoriesList(α *GetRTDataCategoriesList) (β *GetRTDataCategoriesListResponse, err error)

	// ResetTotalValue was auto-generated from WSDL.
	ResetTotalValue(α *ResetTotalValue) (β *ResetTotalValueResponse, err error)

	// GetSites was auto-generated from WSDL.
	GetSites(α *GetSites) (β *GetSitesResponse, err error)

	// SonosAddTrackInPlaylist was auto-generated from WSDL.
	SonosAddTrackInPlaylist(α *SonosAddTrackInPlaylist) (β *SonosAddTrackInPlaylistResponse, err error)

	// GetAlarmFaultObjectList was auto-generated from WSDL.
	GetAlarmFaultObjectList(α *GetAlarmFaultObjectList) (β *GetAlarmFaultObjectListResponse, err error)
}

// CategoryValueType was auto-generated from WSDL.
type categoryValueType string

// Validate validates categoryValueType.
func (v categoryValueType) Validate() bool {
	for _, vv := range []string{
		"total",
		"heating",
		"cooling",
		"domestic_hot_water",
		"plug",
		"undetermined",
		"hot",
		"cold",
		"indoor",
		"outdoor",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// ChartType was auto-generated from WSDL.
type chartType string

// Validate validates chartType.
func (v chartType) Validate() bool {
	for _, vv := range []string{
		"hour",
		"day",
		"week",
		"month",
		"year",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// InstantType was auto-generated from WSDL.
type instantType string

// Validate validates instantType.
func (v instantType) Validate() bool {
	for _, vv := range []string{
		"CURRENT_ELEC",
		"POWER_ELEC",
		"POWER_GAZ",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// MobileUploadType was auto-generated from WSDL.
type mobileUploadType string

// Validate validates mobileUploadType.
func (v mobileUploadType) Validate() bool {
	for _, vv := range []string{
		"MOBILE",
		"PROJECT",
		"ALL_PROJECTS",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// Month was auto-generated from WSDL.
type month string

// Validate validates month.
func (v month) Validate() bool {
	for _, vv := range []string{
		"JANUARY",
		"FEBRUARY",
		"MARCH",
		"APRIL",
		"MAY",
		"JUNE",
		"JULY",
		"AUGUST",
		"SEPTEMBER",
		"OCTOBER",
		"NOVEMBER",
		"DECEMBER",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// RtDataType was auto-generated from WSDL.
type rtDataType string

// Validate validates rtDataType.
func (v rtDataType) Validate() bool {
	for _, vv := range []string{
		"elec",
		"gas",
		"water",
		"undef",
		"temp",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// WsErrorCode was auto-generated from WSDL.
type wsErrorCode string

// Validate validates wsErrorCode.
func (v wsErrorCode) Validate() bool {
	for _, vv := range []string{
		"SERVER_ERROR",
		"USAGE_ERROR",
		"USER_ERROR",
		"AUTHORIZATION_BY_PASS",
		"UNAUTHORIZED",
		"ADMIN_ONLY",
		"NOT_STARTED",
		"DEVICE_DOESNT_EXIST",
		"INVALID_SESSION",
		"INVALID_PASSWORD",
		"TOO_MANY_TRIES",
		"DISABLED_MODULE",
		"SERVER_NOT_ENABLED",
		"INVALID_STATE",
		"INCOMPATIBLE_APPLICATION",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// XxdAccessLevel was auto-generated from WSDL.
type xxdAccessLevel string

// Validate validates xxdAccessLevel.
func (v xxdAccessLevel) Validate() bool {
	for _, vv := range []string{
		"MASTER",
		"USER",
		"LIMITED",
		"DENIED",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// XxdAlarmFaultCategory was auto-generated from WSDL.
type xxdAlarmFaultCategory string

// Validate validates xxdAlarmFaultCategory.
func (v xxdAlarmFaultCategory) Validate() bool {
	for _, vv := range []string{
		"OPEN_ENTRANCE",
		"SELF_PROTECTIONS",
		"OTHERS",
		"INHIBITIONS",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// XxdAlarmHistoryCategory was auto-generated from WSDL.
type xxdAlarmHistoryCategory string

// Validate validates xxdAlarmHistoryCategory.
func (v xxdAlarmHistoryCategory) Validate() bool {
	for _, vv := range []string{
		"ALL",
		"ON_OFF",
		"WITHOUT_ON_OFF",
		"UNACKNOWLEDGED",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// ApplyMobilePicture was auto-generated from WSDL.
type ApplyMobilePicture struct {
	XMLName     xml.Name         `xml:"http://Mobile.ws.domoboxbusiness.com/ ApplyMobilePicture" json:"-" yaml:"-"`
	Session_key string           `xml:"session_key,omitempty" json:"session_key,omitempty" yaml:"session_key,omitempty"`
	Picture_key string           `xml:"picture_key,omitempty" json:"picture_key,omitempty" yaml:"picture_key,omitempty"`
	Room_key    string           `xml:"room_key,omitempty" json:"room_key,omitempty" yaml:"room_key,omitempty"`
	Type        mobileUploadType `xml:"type,omitempty" json:"type,omitempty" yaml:"type,omitempty"`
	Password    string           `xml:"password,omitempty" json:"password,omitempty" yaml:"password,omitempty"`
}

// CancelMobilePicture was auto-generated from WSDL.
type CancelMobilePicture struct {
	XMLName     xml.Name `xml:"http://Mobile.ws.domoboxbusiness.com/ CancelMobilePicture" json:"-" yaml:"-"`
	Session_key string   `xml:"session_key,omitempty" json:"session_key,omitempty" yaml:"session_key,omitempty"`
	Room_key    string   `xml:"room_key,omitempty" json:"room_key,omitempty" yaml:"room_key,omitempty"`
}

// ClearAndroidNotificationDataBase was auto-generated from WSDL.
type ClearAndroidNotificationDataBase struct {
	XMLName     xml.Name `xml:"http://Mobile.ws.domoboxbusiness.com/ ClearAndroidNotificationDataBase" json:"-" yaml:"-"`
	Session_key string   `xml:"session_key,omitempty" json:"session_key,omitempty" yaml:"session_key,omitempty"`
}

// ExecuteAction was auto-generated from WSDL.
type ExecuteAction struct {
	XMLName      xml.Name `xml:"http://Mobile.ws.domoboxbusiness.com/ ExecuteAction" json:"-" yaml:"-"`
	Session_key  string   `xml:"session_key,omitempty" json:"session_key,omitempty" yaml:"session_key,omitempty"`
	Target_key   string   `xml:"target_key,omitempty" json:"target_key,omitempty" yaml:"target_key,omitempty"`
	Prop_clsid   string   `xml:"prop_clsid,omitempty" json:"prop_clsid,omitempty" yaml:"prop_clsid,omitempty"`
	Action_clsid string   `xml:"action_clsid,omitempty" json:"action_clsid,omitempty" yaml:"action_clsid,omitempty"`
	Prop_numr    int      `xml:"prop_numr,omitempty" json:"prop_numr,omitempty" yaml:"prop_numr,omitempty"`
	Descriptor   string   `xml:"descriptor,omitempty" json:"descriptor,omitempty" yaml:"descriptor,omitempty"`
	Password     string   `xml:"password,omitempty" json:"password,omitempty" yaml:"password,omitempty"`
}

// ExecuteActionResponse was auto-generated from WSDL.
type ExecuteActionResponse struct {
	Return bool `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// GetActionTranslationsResponse was auto-generated from WSDL.
type GetActionTranslationsResponse struct {
	Return string `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// GetAlarmFaultObjectList was auto-generated from WSDL.
type GetAlarmFaultObjectList struct {
	XMLName     xml.Name              `xml:"http://Mobile.ws.domoboxbusiness.com/ GetAlarmFaultObjectList" json:"-" yaml:"-"`
	Session_key string                `xml:"session_key,omitempty" json:"session_key,omitempty" yaml:"session_key,omitempty"`
	Site_key    string                `xml:"site_key,omitempty" json:"site_key,omitempty" yaml:"site_key,omitempty"`
	Device_key  string                `xml:"device_key,omitempty" json:"device_key,omitempty" yaml:"device_key,omitempty"`
	Access_code string                `xml:"access_code,omitempty" json:"access_code,omitempty" yaml:"access_code,omitempty"`
	Category    xxdAlarmFaultCategory `xml:"category,omitempty" json:"category,omitempty" yaml:"category,omitempty"`
	Count       int                   `xml:"count,omitempty" json:"count,omitempty" yaml:"count,omitempty"`
	RowIndex    int                   `xml:"rowIndex,omitempty" json:"rowIndex,omitempty" yaml:"rowIndex,omitempty"`
}

// GetAlarmFaultObjectListResponse was auto-generated from WSDL.
type GetAlarmFaultObjectListResponse struct {
	Return []*XxdAlarmFaultDescriptor `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// GetAlarmFaultsStatistics was auto-generated from WSDL.
type GetAlarmFaultsStatistics struct {
	XMLName     xml.Name `xml:"http://Mobile.ws.domoboxbusiness.com/ GetAlarmFaultsStatistics" json:"-" yaml:"-"`
	Session_key string   `xml:"session_key,omitempty" json:"session_key,omitempty" yaml:"session_key,omitempty"`
	Site_key    string   `xml:"site_key,omitempty" json:"site_key,omitempty" yaml:"site_key,omitempty"`
	Device_key  string   `xml:"device_key,omitempty" json:"device_key,omitempty" yaml:"device_key,omitempty"`
	Access_code string   `xml:"access_code,omitempty" json:"access_code,omitempty" yaml:"access_code,omitempty"`
}

// GetAlarmFaultsStatisticsResponse was auto-generated from WSDL.
type GetAlarmFaultsStatisticsResponse struct {
	Return *XxdAlarmFaultsStatisticsDescriptor `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// GetAlarmHistoryObjectList was auto-generated from WSDL.
type GetAlarmHistoryObjectList struct {
	XMLName     xml.Name                `xml:"http://Mobile.ws.domoboxbusiness.com/ GetAlarmHistoryObjectList" json:"-" yaml:"-"`
	Session_key string                  `xml:"session_key,omitempty" json:"session_key,omitempty" yaml:"session_key,omitempty"`
	Site_key    string                  `xml:"site_key,omitempty" json:"site_key,omitempty" yaml:"site_key,omitempty"`
	Device_key  string                  `xml:"device_key,omitempty" json:"device_key,omitempty" yaml:"device_key,omitempty"`
	Access_code string                  `xml:"access_code,omitempty" json:"access_code,omitempty" yaml:"access_code,omitempty"`
	Category    xxdAlarmHistoryCategory `xml:"category,omitempty" json:"category,omitempty" yaml:"category,omitempty"`
	Count       int                     `xml:"count,omitempty" json:"count,omitempty" yaml:"count,omitempty"`
	RowIndex    int                     `xml:"rowIndex,omitempty" json:"rowIndex,omitempty" yaml:"rowIndex,omitempty"`
}

// GetAlarmHistoryObjectListResponse was auto-generated from WSDL.
type GetAlarmHistoryObjectListResponse struct {
	Return []*XxdAlarmHistory `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// GetAlarmeAccessLevel was auto-generated from WSDL.
type GetAlarmeAccessLevel struct {
	XMLName     xml.Name `xml:"http://Mobile.ws.domoboxbusiness.com/ GetAlarmeAccessLevel" json:"-" yaml:"-"`
	Session_key string   `xml:"session_key,omitempty" json:"session_key,omitempty" yaml:"session_key,omitempty"`
	Site_key    string   `xml:"site_key,omitempty" json:"site_key,omitempty" yaml:"site_key,omitempty"`
	Device_key  string   `xml:"device_key,omitempty" json:"device_key,omitempty" yaml:"device_key,omitempty"`
	Access_code string   `xml:"access_code,omitempty" json:"access_code,omitempty" yaml:"access_code,omitempty"`
}

// GetAlarmeAccessLevelResponse was auto-generated from WSDL.
type GetAlarmeAccessLevelResponse struct {
	Return xxdAccessLevel `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// GetAlertLogs was auto-generated from WSDL.
type GetAlertLogs struct {
	XMLName     xml.Name `xml:"http://Mobile.ws.domoboxbusiness.com/ GetAlertLogs" json:"-" yaml:"-"`
	Session_key string   `xml:"session_key,omitempty" json:"session_key,omitempty" yaml:"session_key,omitempty"`
	Start_index int      `xml:"start_index,omitempty" json:"start_index,omitempty" yaml:"start_index,omitempty"`
	Count       int      `xml:"count,omitempty" json:"count,omitempty" yaml:"count,omitempty"`
}

// GetAlertLogsResponse was auto-generated from WSDL.
type GetAlertLogsResponse struct {
	Return *AlertLogs `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// GetBookmarks was auto-generated from WSDL.
type GetBookmarks struct {
	XMLName       xml.Name `xml:"http://Mobile.ws.domoboxbusiness.com/ GetBookmarks" json:"-" yaml:"-"`
	Session_key   string   `xml:"session_key,omitempty" json:"session_key,omitempty" yaml:"session_key,omitempty"`
	Device_list   string   `xml:"device_list,omitempty" json:"device_list,omitempty" yaml:"device_list,omitempty"`
	Group_list    string   `xml:"group_list,omitempty" json:"group_list,omitempty" yaml:"group_list,omitempty"`
	Scenario_list string   `xml:"scenario_list,omitempty" json:"scenario_list,omitempty" yaml:"scenario_list,omitempty"`
}

// GetBookmarksResponse was auto-generated from WSDL.
type GetBookmarksResponse struct {
	Return *Bookmarks `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// GetCameras was auto-generated from WSDL.
type GetCameras struct {
	XMLName     xml.Name `xml:"http://Mobile.ws.domoboxbusiness.com/ GetCameras" json:"-" yaml:"-"`
	Session_key string   `xml:"session_key,omitempty" json:"session_key,omitempty" yaml:"session_key,omitempty"`
	All         bool     `xml:"all,omitempty" json:"all,omitempty" yaml:"all,omitempty"`
}

// GetCamerasResponse was auto-generated from WSDL.
type GetCamerasResponse struct {
	Return *Cameras `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// GetCategories was auto-generated from WSDL.
type GetCategories struct {
	XMLName     xml.Name `xml:"http://Mobile.ws.domoboxbusiness.com/ GetCategories" json:"-" yaml:"-"`
	Session_key string   `xml:"session_key,omitempty" json:"session_key,omitempty" yaml:"session_key,omitempty"`
	Room_key    string   `xml:"room_key,omitempty" json:"room_key,omitempty" yaml:"room_key,omitempty"`
}

// GetCategoriesResponse was auto-generated from WSDL.
type GetCategoriesResponse struct {
	Return *Categories `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// GetCategoriesTranslationsResponse was auto-generated from WSDL.
type GetCategoriesTranslationsResponse struct {
	Return string `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// GetConsumptions was auto-generated from WSDL.
type GetConsumptions struct {
	XMLName     xml.Name `xml:"http://Mobile.ws.domoboxbusiness.com/ GetConsumptions" json:"-" yaml:"-"`
	Session_key string   `xml:"session_key,omitempty" json:"session_key,omitempty" yaml:"session_key,omitempty"`
}

// GetConsumptionsResponse was auto-generated from WSDL.
type GetConsumptionsResponse struct {
	Return *Consumptions `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// GetDeviceInstantane was auto-generated from WSDL.
type GetDeviceInstantane struct {
	XMLName      xml.Name    `xml:"http://Mobile.ws.domoboxbusiness.com/ GetDeviceInstantane" json:"-" yaml:"-"`
	Session_key  string      `xml:"session_key,omitempty" json:"session_key,omitempty" yaml:"session_key,omitempty"`
	Site_key     string      `xml:"site_key,omitempty" json:"site_key,omitempty" yaml:"site_key,omitempty"`
	Device_key   string      `xml:"device_key,omitempty" json:"device_key,omitempty" yaml:"device_key,omitempty"`
	Instant_type instantType `xml:"instant_type,omitempty" json:"instant_type,omitempty" yaml:"instant_type,omitempty"`
}

// GetDeviceInstantaneResponse was auto-generated from WSDL.
type GetDeviceInstantaneResponse struct {
	Return *InstantValue `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// GetDeviceState was auto-generated from WSDL.
type GetDeviceState struct {
	XMLName     xml.Name `xml:"http://Mobile.ws.domoboxbusiness.com/ GetDeviceState" json:"-" yaml:"-"`
	Session_key string   `xml:"session_key,omitempty" json:"session_key,omitempty" yaml:"session_key,omitempty"`
	Device_key  string   `xml:"device_key,omitempty" json:"device_key,omitempty" yaml:"device_key,omitempty"`
}

// GetDeviceStateResponse was auto-generated from WSDL.
type GetDeviceStateResponse struct {
	Return *Device `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// GetDeviceTranslationsResponse was auto-generated from WSDL.
type GetDeviceTranslationsResponse struct {
	Return string `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// GetDevices was auto-generated from WSDL.
type GetDevices struct {
	XMLName        xml.Name `xml:"http://Mobile.ws.domoboxbusiness.com/ GetDevices" json:"-" yaml:"-"`
	Session_key    string   `xml:"session_key,omitempty" json:"session_key,omitempty" yaml:"session_key,omitempty"`
	Room_key       string   `xml:"room_key,omitempty" json:"room_key,omitempty" yaml:"room_key,omitempty"`
	Category_clsid string   `xml:"category_clsid,omitempty" json:"category_clsid,omitempty" yaml:"category_clsid,omitempty"`
}

// GetDevicesFromCatg was auto-generated from WSDL.
type GetDevicesFromCatg struct {
	XMLName        xml.Name `xml:"http://Mobile.ws.domoboxbusiness.com/ GetDevicesFromCatg" json:"-" yaml:"-"`
	Session_key    string   `xml:"session_key,omitempty" json:"session_key,omitempty" yaml:"session_key,omitempty"`
	Category_clsid string   `xml:"category_clsid,omitempty" json:"category_clsid,omitempty" yaml:"category_clsid,omitempty"`
}

// GetDevicesFromCatgResponse was auto-generated from WSDL.
type GetDevicesFromCatgResponse struct {
	Return *Devices `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// GetDevicesResponse was auto-generated from WSDL.
type GetDevicesResponse struct {
	Return *Devices `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// GetGroup was auto-generated from WSDL.
type GetGroup struct {
	XMLName     xml.Name `xml:"http://Mobile.ws.domoboxbusiness.com/ GetGroup" json:"-" yaml:"-"`
	Session_key string   `xml:"session_key,omitempty" json:"session_key,omitempty" yaml:"session_key,omitempty"`
	Group_key   string   `xml:"group_key,omitempty" json:"group_key,omitempty" yaml:"group_key,omitempty"`
}

// GetGroupResponse was auto-generated from WSDL.
type GetGroupResponse struct {
	Return *Group `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// GetGroups was auto-generated from WSDL.
type GetGroups struct {
	XMLName     xml.Name `xml:"http://Mobile.ws.domoboxbusiness.com/ GetGroups" json:"-" yaml:"-"`
	Session_key string   `xml:"session_key,omitempty" json:"session_key,omitempty" yaml:"session_key,omitempty"`
}

// GetGroupsResponse was auto-generated from WSDL.
type GetGroupsResponse struct {
	Return *Groups `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// GetHistValue was auto-generated from WSDL.
type GetHistValue struct {
	XMLName     xml.Name          `xml:"http://Mobile.ws.domoboxbusiness.com/ GetHistValue" json:"-" yaml:"-"`
	Session_key string            `xml:"session_key,omitempty" json:"session_key,omitempty" yaml:"session_key,omitempty"`
	Site_key    string            `xml:"site_key,omitempty" json:"site_key,omitempty" yaml:"site_key,omitempty"`
	Device_key  string            `xml:"device_key,omitempty" json:"device_key,omitempty" yaml:"device_key,omitempty"`
	Chart_type  chartType         `xml:"chart_type,omitempty" json:"chart_type,omitempty" yaml:"chart_type,omitempty"`
	Type        rtDataType        `xml:"type,omitempty" json:"type,omitempty" yaml:"type,omitempty"`
	Category    categoryValueType `xml:"category,omitempty" json:"category,omitempty" yaml:"category,omitempty"`
	Date        *LdDate           `xml:"date,omitempty" json:"date,omitempty" yaml:"date,omitempty"`
}

// GetHistValueResponse was auto-generated from WSDL.
type GetHistValueResponse struct {
	Return *HistogramValue `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// GetInstantTypeList was auto-generated from WSDL.
type GetInstantTypeList struct {
	XMLName     xml.Name `xml:"http://Mobile.ws.domoboxbusiness.com/ GetInstantTypeList" json:"-" yaml:"-"`
	Session_key string   `xml:"session_key,omitempty" json:"session_key,omitempty" yaml:"session_key,omitempty"`
	Site_key    string   `xml:"site_key,omitempty" json:"site_key,omitempty" yaml:"site_key,omitempty"`
	Device_key  string   `xml:"device_key,omitempty" json:"device_key,omitempty" yaml:"device_key,omitempty"`
}

// GetInstantTypeListResponse was auto-generated from WSDL.
type GetInstantTypeListResponse struct {
	Return []instantType `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// GetPicture was auto-generated from WSDL.
type GetPicture struct {
	XMLName     xml.Name `xml:"http://Mobile.ws.domoboxbusiness.com/ GetPicture" json:"-" yaml:"-"`
	Session_key string   `xml:"session_key,omitempty" json:"session_key,omitempty" yaml:"session_key,omitempty"`
	Picture_key string   `xml:"picture_key,omitempty" json:"picture_key,omitempty" yaml:"picture_key,omitempty"`
}

// GetPictureResponse was auto-generated from WSDL.
type GetPictureResponse struct {
	Return []byte `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// GetPieChartValue was auto-generated from WSDL.
type GetPieChartValue struct {
	XMLName     xml.Name   `xml:"http://Mobile.ws.domoboxbusiness.com/ GetPieChartValue" json:"-" yaml:"-"`
	Session_key string     `xml:"session_key,omitempty" json:"session_key,omitempty" yaml:"session_key,omitempty"`
	Site_key    string     `xml:"site_key,omitempty" json:"site_key,omitempty" yaml:"site_key,omitempty"`
	Device_key  string     `xml:"device_key,omitempty" json:"device_key,omitempty" yaml:"device_key,omitempty"`
	Chart_type  chartType  `xml:"chart_type,omitempty" json:"chart_type,omitempty" yaml:"chart_type,omitempty"`
	Type        rtDataType `xml:"type,omitempty" json:"type,omitempty" yaml:"type,omitempty"`
	Date        *LdDate    `xml:"date,omitempty" json:"date,omitempty" yaml:"date,omitempty"`
}

// GetPieChartValueResponse was auto-generated from WSDL.
type GetPieChartValueResponse struct {
	Return *PieChartValue `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// GetRTDataCategoriesList was auto-generated from WSDL.
type GetRTDataCategoriesList struct {
	XMLName     xml.Name   `xml:"http://Mobile.ws.domoboxbusiness.com/ GetRTDataCategoriesList" json:"-" yaml:"-"`
	Session_key string     `xml:"session_key,omitempty" json:"session_key,omitempty" yaml:"session_key,omitempty"`
	Site_key    string     `xml:"site_key,omitempty" json:"site_key,omitempty" yaml:"site_key,omitempty"`
	Device_key  string     `xml:"device_key,omitempty" json:"device_key,omitempty" yaml:"device_key,omitempty"`
	Type        rtDataType `xml:"type,omitempty" json:"type,omitempty" yaml:"type,omitempty"`
}

// GetRTDataCategoriesListResponse was auto-generated from WSDL.
type GetRTDataCategoriesListResponse struct {
	Return []categoryValueType `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// GetRTDataTypeList was auto-generated from WSDL.
type GetRTDataTypeList struct {
	XMLName     xml.Name `xml:"http://Mobile.ws.domoboxbusiness.com/ GetRTDataTypeList" json:"-" yaml:"-"`
	Session_key string   `xml:"session_key,omitempty" json:"session_key,omitempty" yaml:"session_key,omitempty"`
	Site_key    string   `xml:"site_key,omitempty" json:"site_key,omitempty" yaml:"site_key,omitempty"`
	Device_key  string   `xml:"device_key,omitempty" json:"device_key,omitempty" yaml:"device_key,omitempty"`
}

// GetRTDataTypeListResponse was auto-generated from WSDL.
type GetRTDataTypeListResponse struct {
	Return []rtDataType `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// GetRTPageDescriptor was auto-generated from WSDL.
type GetRTPageDescriptor struct {
	XMLName     xml.Name `xml:"http://Mobile.ws.domoboxbusiness.com/ GetRTPageDescriptor" json:"-" yaml:"-"`
	Session_key string   `xml:"session_key,omitempty" json:"session_key,omitempty" yaml:"session_key,omitempty"`
	Site_key    string   `xml:"site_key,omitempty" json:"site_key,omitempty" yaml:"site_key,omitempty"`
	Device_key  string   `xml:"device_key,omitempty" json:"device_key,omitempty" yaml:"device_key,omitempty"`
}

// GetRTPageDescriptorResponse was auto-generated from WSDL.
type GetRTPageDescriptorResponse struct {
	Return *RtPageDescriptor `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// GetRooms was auto-generated from WSDL.
type GetRooms struct {
	XMLName     xml.Name `xml:"http://Mobile.ws.domoboxbusiness.com/ GetRooms" json:"-" yaml:"-"`
	Session_key string   `xml:"session_key,omitempty" json:"session_key,omitempty" yaml:"session_key,omitempty"`
}

// GetRoomsNew was auto-generated from WSDL.
type GetRoomsNew struct {
	XMLName     xml.Name `xml:"http://Mobile.ws.domoboxbusiness.com/ GetRoomsNew" json:"-" yaml:"-"`
	Session_key string   `xml:"session_key,omitempty" json:"session_key,omitempty" yaml:"session_key,omitempty"`
}

// GetRoomsNewResponse was auto-generated from WSDL.
type GetRoomsNewResponse struct {
	Return *Rooms `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// GetRoomsResponse was auto-generated from WSDL.
type GetRoomsResponse struct {
	Return *Rooms `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// GetScenarios was auto-generated from WSDL.
type GetScenarios struct {
	XMLName     xml.Name `xml:"http://Mobile.ws.domoboxbusiness.com/ GetScenarios" json:"-" yaml:"-"`
	Session_key string   `xml:"session_key,omitempty" json:"session_key,omitempty" yaml:"session_key,omitempty"`
}

// GetScenariosResponse was auto-generated from WSDL.
type GetScenariosResponse struct {
	Return *Scenarios `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// GetSitesResponse was auto-generated from WSDL.
type GetSitesResponse struct {
	Return *Sites `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// GetStateTranslationsResponse was auto-generated from WSDL.
type GetStateTranslationsResponse struct {
	Return string `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// GetTotalDataValue was auto-generated from WSDL.
type GetTotalDataValue struct {
	XMLName     xml.Name          `xml:"http://Mobile.ws.domoboxbusiness.com/ GetTotalDataValue" json:"-" yaml:"-"`
	Session_key string            `xml:"session_key,omitempty" json:"session_key,omitempty" yaml:"session_key,omitempty"`
	Site_key    string            `xml:"site_key,omitempty" json:"site_key,omitempty" yaml:"site_key,omitempty"`
	Device_key  string            `xml:"device_key,omitempty" json:"device_key,omitempty" yaml:"device_key,omitempty"`
	Type        rtDataType        `xml:"type,omitempty" json:"type,omitempty" yaml:"type,omitempty"`
	Category    categoryValueType `xml:"category,omitempty" json:"category,omitempty" yaml:"category,omitempty"`
}

// GetTotalDataValueResponse was auto-generated from WSDL.
type GetTotalDataValueResponse struct {
	Return *TotalValue `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// GetTranslationsRESTResponse was auto-generated from WSDL.
type GetTranslationsRESTResponse struct {
	Return []*MobileTranslationDescriptor `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// GetTranslationsResponse was auto-generated from WSDL.
type GetTranslationsResponse struct {
	Return string `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// GetUsers was auto-generated from WSDL.
type GetUsers struct {
	XMLName  xml.Name `xml:"http://Mobile.ws.domoboxbusiness.com/ GetUsers" json:"-" yaml:"-"`
	Site_key string   `xml:"site_key,omitempty" json:"site_key,omitempty" yaml:"site_key,omitempty"`
}

// GetUsersResponse was auto-generated from WSDL.
type GetUsersResponse struct {
	Return *Users `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// GetVarunaAlarms was auto-generated from WSDL.
type GetVarunaAlarms struct {
	XMLName     xml.Name `xml:"http://Mobile.ws.domoboxbusiness.com/ GetVarunaAlarms" json:"-" yaml:"-"`
	Session_key string   `xml:"session_key,omitempty" json:"session_key,omitempty" yaml:"session_key,omitempty"`
	Device_key  string   `xml:"device_key,omitempty" json:"device_key,omitempty" yaml:"device_key,omitempty"`
}

// GetVarunaAlarmsResponse was auto-generated from WSDL.
type GetVarunaAlarmsResponse struct {
	Return []*VarunaInfos `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// GetVarunaEvents was auto-generated from WSDL.
type GetVarunaEvents struct {
	XMLName     xml.Name `xml:"http://Mobile.ws.domoboxbusiness.com/ GetVarunaEvents" json:"-" yaml:"-"`
	Session_key string   `xml:"session_key,omitempty" json:"session_key,omitempty" yaml:"session_key,omitempty"`
	Device_key  string   `xml:"device_key,omitempty" json:"device_key,omitempty" yaml:"device_key,omitempty"`
}

// GetVarunaEventsResponse was auto-generated from WSDL.
type GetVarunaEventsResponse struct {
	Return []*VarunaInfos `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// GetVoiceObjects was auto-generated from WSDL.
type GetVoiceObjects struct {
	XMLName     xml.Name `xml:"http://Mobile.ws.domoboxbusiness.com/ GetVoiceObjects" json:"-" yaml:"-"`
	Session_key string   `xml:"session_key,omitempty" json:"session_key,omitempty" yaml:"session_key,omitempty"`
}

// GetVoiceObjectsResponse was auto-generated from WSDL.
type GetVoiceObjectsResponse struct {
	Return *Voicestruct `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// KeepAlive was auto-generated from WSDL.
type KeepAlive struct {
	XMLName     xml.Name `xml:"http://Mobile.ws.domoboxbusiness.com/ KeepAlive" json:"-" yaml:"-"`
	Session_key string   `xml:"session_key,omitempty" json:"session_key,omitempty" yaml:"session_key,omitempty"`
}

// Login was auto-generated from WSDL.
type Login struct {
	XMLName  xml.Name `xml:"http://Mobile.ws.domoboxbusiness.com/ Login" json:"-" yaml:"-"`
	Site_key string   `xml:"site_key,omitempty" json:"site_key,omitempty" yaml:"site_key,omitempty"`
	User_key string   `xml:"user_key,omitempty" json:"user_key,omitempty" yaml:"user_key,omitempty"`
	Password string   `xml:"password,omitempty" json:"password,omitempty" yaml:"password,omitempty"`
}

// LoginAlertLogs was auto-generated from WSDL.
type LoginAlertLogs struct {
	XMLName  xml.Name `xml:"http://Mobile.ws.domoboxbusiness.com/ LoginAlertLogs" json:"-" yaml:"-"`
	Site_key string   `xml:"site_key,omitempty" json:"site_key,omitempty" yaml:"site_key,omitempty"`
	User_key string   `xml:"user_key,omitempty" json:"user_key,omitempty" yaml:"user_key,omitempty"`
	Password string   `xml:"password,omitempty" json:"password,omitempty" yaml:"password,omitempty"`
}

// LoginAlertLogsResponse was auto-generated from WSDL.
type LoginAlertLogsResponse struct {
	Return *AlertLogs `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// LoginBookmarks was auto-generated from WSDL.
type LoginBookmarks struct {
	XMLName       xml.Name `xml:"http://Mobile.ws.domoboxbusiness.com/ LoginBookmarks" json:"-" yaml:"-"`
	Site_key      string   `xml:"site_key,omitempty" json:"site_key,omitempty" yaml:"site_key,omitempty"`
	User_key      string   `xml:"user_key,omitempty" json:"user_key,omitempty" yaml:"user_key,omitempty"`
	Password      string   `xml:"password,omitempty" json:"password,omitempty" yaml:"password,omitempty"`
	Device_list   string   `xml:"device_list,omitempty" json:"device_list,omitempty" yaml:"device_list,omitempty"`
	Group_list    string   `xml:"group_list,omitempty" json:"group_list,omitempty" yaml:"group_list,omitempty"`
	Scenario_list string   `xml:"scenario_list,omitempty" json:"scenario_list,omitempty" yaml:"scenario_list,omitempty"`
}

// LoginBookmarksResponse was auto-generated from WSDL.
type LoginBookmarksResponse struct {
	Return *Bookmarks `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// LoginCameras was auto-generated from WSDL.
type LoginCameras struct {
	XMLName  xml.Name `xml:"http://Mobile.ws.domoboxbusiness.com/ LoginCameras" json:"-" yaml:"-"`
	Site_key string   `xml:"site_key,omitempty" json:"site_key,omitempty" yaml:"site_key,omitempty"`
	User_key string   `xml:"user_key,omitempty" json:"user_key,omitempty" yaml:"user_key,omitempty"`
	Password string   `xml:"password,omitempty" json:"password,omitempty" yaml:"password,omitempty"`
	All      bool     `xml:"all,omitempty" json:"all,omitempty" yaml:"all,omitempty"`
}

// LoginCamerasResponse was auto-generated from WSDL.
type LoginCamerasResponse struct {
	Return *Cameras `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// LoginConsumptions was auto-generated from WSDL.
type LoginConsumptions struct {
	XMLName  xml.Name `xml:"http://Mobile.ws.domoboxbusiness.com/ LoginConsumptions" json:"-" yaml:"-"`
	Site_key string   `xml:"site_key,omitempty" json:"site_key,omitempty" yaml:"site_key,omitempty"`
	User_key string   `xml:"user_key,omitempty" json:"user_key,omitempty" yaml:"user_key,omitempty"`
	Password string   `xml:"password,omitempty" json:"password,omitempty" yaml:"password,omitempty"`
}

// LoginConsumptionsResponse was auto-generated from WSDL.
type LoginConsumptionsResponse struct {
	Return *Consumptions `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// LoginDeviceState was auto-generated from WSDL.
type LoginDeviceState struct {
	XMLName    xml.Name `xml:"http://Mobile.ws.domoboxbusiness.com/ LoginDeviceState" json:"-" yaml:"-"`
	Site_key   string   `xml:"site_key,omitempty" json:"site_key,omitempty" yaml:"site_key,omitempty"`
	User_key   string   `xml:"user_key,omitempty" json:"user_key,omitempty" yaml:"user_key,omitempty"`
	Password   string   `xml:"password,omitempty" json:"password,omitempty" yaml:"password,omitempty"`
	Device_key string   `xml:"device_key,omitempty" json:"device_key,omitempty" yaml:"device_key,omitempty"`
}

// LoginDeviceStateResponse was auto-generated from WSDL.
type LoginDeviceStateResponse struct {
	Return *Devices `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// LoginDevices was auto-generated from WSDL.
type LoginDevices struct {
	XMLName        xml.Name `xml:"http://Mobile.ws.domoboxbusiness.com/ LoginDevices" json:"-" yaml:"-"`
	Site_key       string   `xml:"site_key,omitempty" json:"site_key,omitempty" yaml:"site_key,omitempty"`
	User_key       string   `xml:"user_key,omitempty" json:"user_key,omitempty" yaml:"user_key,omitempty"`
	Password       string   `xml:"password,omitempty" json:"password,omitempty" yaml:"password,omitempty"`
	Room_key       string   `xml:"room_key,omitempty" json:"room_key,omitempty" yaml:"room_key,omitempty"`
	Category_clsid string   `xml:"category_clsid,omitempty" json:"category_clsid,omitempty" yaml:"category_clsid,omitempty"`
}

// LoginDevicesFromCatg was auto-generated from WSDL.
type LoginDevicesFromCatg struct {
	XMLName        xml.Name `xml:"http://Mobile.ws.domoboxbusiness.com/ LoginDevicesFromCatg" json:"-" yaml:"-"`
	Site_key       string   `xml:"site_key,omitempty" json:"site_key,omitempty" yaml:"site_key,omitempty"`
	User_key       string   `xml:"user_key,omitempty" json:"user_key,omitempty" yaml:"user_key,omitempty"`
	Password       string   `xml:"password,omitempty" json:"password,omitempty" yaml:"password,omitempty"`
	Category_clsid string   `xml:"category_clsid,omitempty" json:"category_clsid,omitempty" yaml:"category_clsid,omitempty"`
}

// LoginDevicesFromCatgResponse was auto-generated from WSDL.
type LoginDevicesFromCatgResponse struct {
	Return *Devices `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// LoginDevicesResponse was auto-generated from WSDL.
type LoginDevicesResponse struct {
	Return *Devices `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// LoginExecuteAction was auto-generated from WSDL.
type LoginExecuteAction struct {
	XMLName      xml.Name `xml:"http://Mobile.ws.domoboxbusiness.com/ LoginExecuteAction" json:"-" yaml:"-"`
	Site_key     string   `xml:"site_key,omitempty" json:"site_key,omitempty" yaml:"site_key,omitempty"`
	User_key     string   `xml:"user_key,omitempty" json:"user_key,omitempty" yaml:"user_key,omitempty"`
	Password     string   `xml:"password,omitempty" json:"password,omitempty" yaml:"password,omitempty"`
	Target_key   string   `xml:"target_key,omitempty" json:"target_key,omitempty" yaml:"target_key,omitempty"`
	Prop_clsid   string   `xml:"prop_clsid,omitempty" json:"prop_clsid,omitempty" yaml:"prop_clsid,omitempty"`
	Action_clsid string   `xml:"action_clsid,omitempty" json:"action_clsid,omitempty" yaml:"action_clsid,omitempty"`
	Prop_numr    int      `xml:"prop_numr,omitempty" json:"prop_numr,omitempty" yaml:"prop_numr,omitempty"`
	Descriptor   string   `xml:"descriptor,omitempty" json:"descriptor,omitempty" yaml:"descriptor,omitempty"`
}

// LoginExecuteActionResponse was auto-generated from WSDL.
type LoginExecuteActionResponse struct {
	Return bool `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// LoginGroup was auto-generated from WSDL.
type LoginGroup struct {
	XMLName   xml.Name `xml:"http://Mobile.ws.domoboxbusiness.com/ LoginGroup" json:"-" yaml:"-"`
	Site_key  string   `xml:"site_key,omitempty" json:"site_key,omitempty" yaml:"site_key,omitempty"`
	User_key  string   `xml:"user_key,omitempty" json:"user_key,omitempty" yaml:"user_key,omitempty"`
	Password  string   `xml:"password,omitempty" json:"password,omitempty" yaml:"password,omitempty"`
	Group_key string   `xml:"group_key,omitempty" json:"group_key,omitempty" yaml:"group_key,omitempty"`
}

// LoginGroupResponse was auto-generated from WSDL.
type LoginGroupResponse struct {
	Return *Groups `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// LoginGroups was auto-generated from WSDL.
type LoginGroups struct {
	XMLName  xml.Name `xml:"http://Mobile.ws.domoboxbusiness.com/ LoginGroups" json:"-" yaml:"-"`
	Site_key string   `xml:"site_key,omitempty" json:"site_key,omitempty" yaml:"site_key,omitempty"`
	User_key string   `xml:"user_key,omitempty" json:"user_key,omitempty" yaml:"user_key,omitempty"`
	Password string   `xml:"password,omitempty" json:"password,omitempty" yaml:"password,omitempty"`
}

// LoginGroupsResponse was auto-generated from WSDL.
type LoginGroupsResponse struct {
	Return *Groups `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// LoginInfos was auto-generated from WSDL.
type LoginInfos struct {
	XMLName       xml.Name `xml:"http://Mobile.ws.domoboxbusiness.com/ LoginInfos" json:"-" yaml:"-"`
	Site_key      string   `xml:"site_key,omitempty" json:"site_key,omitempty" yaml:"site_key,omitempty"`
	User_key      string   `xml:"user_key,omitempty" json:"user_key,omitempty" yaml:"user_key,omitempty"`
	Password      string   `xml:"password,omitempty" json:"password,omitempty" yaml:"password,omitempty"`
	Device_list   string   `xml:"device_list,omitempty" json:"device_list,omitempty" yaml:"device_list,omitempty"`
	Group_list    string   `xml:"group_list,omitempty" json:"group_list,omitempty" yaml:"group_list,omitempty"`
	Scenario_list string   `xml:"scenario_list,omitempty" json:"scenario_list,omitempty" yaml:"scenario_list,omitempty"`
}

// LoginInfosResponse was auto-generated from WSDL.
type LoginInfosResponse struct {
	Return *Infos `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// LoginResponse was auto-generated from WSDL.
type LoginResponse struct {
	Return string `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// LoginScenarios was auto-generated from WSDL.
type LoginScenarios struct {
	XMLName  xml.Name `xml:"http://Mobile.ws.domoboxbusiness.com/ LoginScenarios" json:"-" yaml:"-"`
	Site_key string   `xml:"site_key,omitempty" json:"site_key,omitempty" yaml:"site_key,omitempty"`
	User_key string   `xml:"user_key,omitempty" json:"user_key,omitempty" yaml:"user_key,omitempty"`
	Password string   `xml:"password,omitempty" json:"password,omitempty" yaml:"password,omitempty"`
}

// LoginScenariosResponse was auto-generated from WSDL.
type LoginScenariosResponse struct {
	Return *Scenarios `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// NewLoginInfos was auto-generated from WSDL.
type NewLoginInfos struct {
	XMLName       xml.Name `xml:"http://Mobile.ws.domoboxbusiness.com/ NewLoginInfos" json:"-" yaml:"-"`
	Site_key      string   `xml:"site_key,omitempty" json:"site_key,omitempty" yaml:"site_key,omitempty"`
	User_key      string   `xml:"user_key,omitempty" json:"user_key,omitempty" yaml:"user_key,omitempty"`
	Password      string   `xml:"password,omitempty" json:"password,omitempty" yaml:"password,omitempty"`
	Device_list   string   `xml:"device_list,omitempty" json:"device_list,omitempty" yaml:"device_list,omitempty"`
	Group_list    string   `xml:"group_list,omitempty" json:"group_list,omitempty" yaml:"group_list,omitempty"`
	Scenario_list string   `xml:"scenario_list,omitempty" json:"scenario_list,omitempty" yaml:"scenario_list,omitempty"`
}

// NewLoginInfosResponse was auto-generated from WSDL.
type NewLoginInfosResponse struct {
	Return *Infos `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// RefreshVarunaEvents was auto-generated from WSDL.
type RefreshVarunaEvents struct {
	XMLName     xml.Name `xml:"http://Mobile.ws.domoboxbusiness.com/ RefreshVarunaEvents" json:"-" yaml:"-"`
	Session_key string   `xml:"session_key,omitempty" json:"session_key,omitempty" yaml:"session_key,omitempty"`
	Device_key  string   `xml:"device_key,omitempty" json:"device_key,omitempty" yaml:"device_key,omitempty"`
}

// RegisterNotif was auto-generated from WSDL.
type RegisterNotif struct {
	XMLName     xml.Name `xml:"http://Mobile.ws.domoboxbusiness.com/ RegisterNotif" json:"-" yaml:"-"`
	Session_key string   `xml:"session_key,omitempty" json:"session_key,omitempty" yaml:"session_key,omitempty"`
	Phone_id    string   `xml:"phone_id,omitempty" json:"phone_id,omitempty" yaml:"phone_id,omitempty"`
}

// RegisterNotifIpad was auto-generated from WSDL.
type RegisterNotifIpad struct {
	XMLName     xml.Name `xml:"http://Mobile.ws.domoboxbusiness.com/ RegisterNotifIpad" json:"-" yaml:"-"`
	Session_key string   `xml:"session_key,omitempty" json:"session_key,omitempty" yaml:"session_key,omitempty"`
	Phone_id    string   `xml:"phone_id,omitempty" json:"phone_id,omitempty" yaml:"phone_id,omitempty"`
}

// RegisterNotifIpad2 was auto-generated from WSDL.
type RegisterNotifIpad2 struct {
	XMLName     xml.Name `xml:"http://Mobile.ws.domoboxbusiness.com/ RegisterNotifIpad2" json:"-" yaml:"-"`
	Session_key string   `xml:"session_key,omitempty" json:"session_key,omitempty" yaml:"session_key,omitempty"`
	Phone_id    string   `xml:"phone_id,omitempty" json:"phone_id,omitempty" yaml:"phone_id,omitempty"`
}

// RegisterNotifPhoneAndroid was auto-generated from WSDL.
type RegisterNotifPhoneAndroid struct {
	XMLName         xml.Name `xml:"http://Mobile.ws.domoboxbusiness.com/ RegisterNotifPhoneAndroid" json:"-" yaml:"-"`
	Session_key     string   `xml:"session_key,omitempty" json:"session_key,omitempty" yaml:"session_key,omitempty"`
	Registration_id string   `xml:"registration_id,omitempty" json:"registration_id,omitempty" yaml:"registration_id,omitempty"`
}

// RegisterNotifPhoneAndroidResponse was auto-generated from WSDL.
type RegisterNotifPhoneAndroidResponse struct {
	Return bool `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// RegisterNotifTabletteAndroid was auto-generated from WSDL.
type RegisterNotifTabletteAndroid struct {
	XMLName         xml.Name `xml:"http://Mobile.ws.domoboxbusiness.com/ RegisterNotifTabletteAndroid" json:"-" yaml:"-"`
	Session_key     string   `xml:"session_key,omitempty" json:"session_key,omitempty" yaml:"session_key,omitempty"`
	Registration_id string   `xml:"registration_id,omitempty" json:"registration_id,omitempty" yaml:"registration_id,omitempty"`
}

// RegisterNotifTabletteAndroidResponse was auto-generated from
// WSDL.
type RegisterNotifTabletteAndroidResponse struct {
	Return bool `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// ResetInstantValue was auto-generated from WSDL.
type ResetInstantValue struct {
	XMLName     xml.Name    `xml:"http://Mobile.ws.domoboxbusiness.com/ ResetInstantValue" json:"-" yaml:"-"`
	Session_key string      `xml:"session_key,omitempty" json:"session_key,omitempty" yaml:"session_key,omitempty"`
	Site_key    string      `xml:"site_key,omitempty" json:"site_key,omitempty" yaml:"site_key,omitempty"`
	Device_key  string      `xml:"device_key,omitempty" json:"device_key,omitempty" yaml:"device_key,omitempty"`
	Type        instantType `xml:"type,omitempty" json:"type,omitempty" yaml:"type,omitempty"`
}

// ResetNotif was auto-generated from WSDL.
type ResetNotif struct {
	XMLName     xml.Name `xml:"http://Mobile.ws.domoboxbusiness.com/ ResetNotif" json:"-" yaml:"-"`
	Session_key string   `xml:"session_key,omitempty" json:"session_key,omitempty" yaml:"session_key,omitempty"`
	Phone_id    string   `xml:"phone_id,omitempty" json:"phone_id,omitempty" yaml:"phone_id,omitempty"`
}

// ResetTotalValue was auto-generated from WSDL.
type ResetTotalValue struct {
	XMLName     xml.Name   `xml:"http://Mobile.ws.domoboxbusiness.com/ ResetTotalValue" json:"-" yaml:"-"`
	Session_key string     `xml:"session_key,omitempty" json:"session_key,omitempty" yaml:"session_key,omitempty"`
	Site_key    string     `xml:"site_key,omitempty" json:"site_key,omitempty" yaml:"site_key,omitempty"`
	Device_key  string     `xml:"device_key,omitempty" json:"device_key,omitempty" yaml:"device_key,omitempty"`
	Type        rtDataType `xml:"type,omitempty" json:"type,omitempty" yaml:"type,omitempty"`
}

// SonosAddAndPlay was auto-generated from WSDL.
type SonosAddAndPlay struct {
	XMLName     xml.Name `xml:"http://Mobile.ws.domoboxbusiness.com/ SonosAddAndPlay" json:"-" yaml:"-"`
	Session_key string   `xml:"session_key,omitempty" json:"session_key,omitempty" yaml:"session_key,omitempty"`
	Device_key  string   `xml:"device_key,omitempty" json:"device_key,omitempty" yaml:"device_key,omitempty"`
	Rincon_uri  string   `xml:"rincon_uri,omitempty" json:"rincon_uri,omitempty" yaml:"rincon_uri,omitempty"`
	Metadata    string   `xml:"metadata,omitempty" json:"metadata,omitempty" yaml:"metadata,omitempty"`
	Num         string   `xml:"num,omitempty" json:"num,omitempty" yaml:"num,omitempty"`
}

// SonosAddFavorite was auto-generated from WSDL.
type SonosAddFavorite struct {
	XMLName     xml.Name `xml:"http://Mobile.ws.domoboxbusiness.com/ SonosAddFavorite" json:"-" yaml:"-"`
	Session_key string   `xml:"session_key,omitempty" json:"session_key,omitempty" yaml:"session_key,omitempty"`
	Device_key  string   `xml:"device_key,omitempty" json:"device_key,omitempty" yaml:"device_key,omitempty"`
	Path        string   `xml:"path,omitempty" json:"path,omitempty" yaml:"path,omitempty"`
}

// SonosAddFavoriteResponse was auto-generated from WSDL.
type SonosAddFavoriteResponse struct {
	Return bool `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// SonosAddTrackInPlaylist was auto-generated from WSDL.
type SonosAddTrackInPlaylist struct {
	XMLName     xml.Name `xml:"http://Mobile.ws.domoboxbusiness.com/ SonosAddTrackInPlaylist" json:"-" yaml:"-"`
	Session_key string   `xml:"session_key,omitempty" json:"session_key,omitempty" yaml:"session_key,omitempty"`
	Device_key  string   `xml:"device_key,omitempty" json:"device_key,omitempty" yaml:"device_key,omitempty"`
	Rincon_uri  string   `xml:"rincon_uri,omitempty" json:"rincon_uri,omitempty" yaml:"rincon_uri,omitempty"`
	Metadata    string   `xml:"metadata,omitempty" json:"metadata,omitempty" yaml:"metadata,omitempty"`
	Num         string   `xml:"num,omitempty" json:"num,omitempty" yaml:"num,omitempty"`
}

// SonosDeleteFavorite was auto-generated from WSDL.
type SonosDeleteFavorite struct {
	XMLName     xml.Name `xml:"http://Mobile.ws.domoboxbusiness.com/ SonosDeleteFavorite" json:"-" yaml:"-"`
	Session_key string   `xml:"session_key,omitempty" json:"session_key,omitempty" yaml:"session_key,omitempty"`
	Device_key  string   `xml:"device_key,omitempty" json:"device_key,omitempty" yaml:"device_key,omitempty"`
	Path        string   `xml:"path,omitempty" json:"path,omitempty" yaml:"path,omitempty"`
}

// SonosDeleteFavoriteResponse was auto-generated from WSDL.
type SonosDeleteFavoriteResponse struct {
	Return bool `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// SonosDeleteList was auto-generated from WSDL.
type SonosDeleteList struct {
	XMLName     xml.Name `xml:"http://Mobile.ws.domoboxbusiness.com/ SonosDeleteList" json:"-" yaml:"-"`
	Session_key string   `xml:"session_key,omitempty" json:"session_key,omitempty" yaml:"session_key,omitempty"`
	Device_key  string   `xml:"device_key,omitempty" json:"device_key,omitempty" yaml:"device_key,omitempty"`
}

// SonosExplore was auto-generated from WSDL.
type SonosExplore struct {
	XMLName        xml.Name `xml:"http://Mobile.ws.domoboxbusiness.com/ SonosExplore" json:"-" yaml:"-"`
	Session_key    string   `xml:"session_key,omitempty" json:"session_key,omitempty" yaml:"session_key,omitempty"`
	Device_key     string   `xml:"device_key,omitempty" json:"device_key,omitempty" yaml:"device_key,omitempty"`
	Path           string   `xml:"path,omitempty" json:"path,omitempty" yaml:"path,omitempty"`
	Starting_index int      `xml:"starting_index,omitempty" json:"starting_index,omitempty" yaml:"starting_index,omitempty"`
	Count          int      `xml:"count,omitempty" json:"count,omitempty" yaml:"count,omitempty"`
}

// SonosExploreResponse was auto-generated from WSDL.
type SonosExploreResponse struct {
	Return []*SonosTrackDescriptor `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// SonosGetAllGroups was auto-generated from WSDL.
type SonosGetAllGroups struct {
	XMLName     xml.Name `xml:"http://Mobile.ws.domoboxbusiness.com/ SonosGetAllGroups" json:"-" yaml:"-"`
	Session_key string   `xml:"session_key,omitempty" json:"session_key,omitempty" yaml:"session_key,omitempty"`
}

// SonosGetAllGroupsResponse was auto-generated from WSDL.
type SonosGetAllGroupsResponse struct {
	Return []*SonosMobileGroupDescriptor `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// SonosGetLineInList was auto-generated from WSDL.
type SonosGetLineInList struct {
	XMLName     xml.Name `xml:"http://Mobile.ws.domoboxbusiness.com/ SonosGetLineInList" json:"-" yaml:"-"`
	Session_key string   `xml:"session_key,omitempty" json:"session_key,omitempty" yaml:"session_key,omitempty"`
}

// SonosGetLineInListResponse was auto-generated from WSDL.
type SonosGetLineInListResponse struct {
	Return []*SonosGroupDescriptor `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// SonosGetPlaylist was auto-generated from WSDL.
type SonosGetPlaylist struct {
	XMLName     xml.Name `xml:"http://Mobile.ws.domoboxbusiness.com/ SonosGetPlaylist" json:"-" yaml:"-"`
	Session_key string   `xml:"session_key,omitempty" json:"session_key,omitempty" yaml:"session_key,omitempty"`
	Device_key  string   `xml:"device_key,omitempty" json:"device_key,omitempty" yaml:"device_key,omitempty"`
}

// SonosGetPlaylistResponse was auto-generated from WSDL.
type SonosGetPlaylistResponse struct {
	Return []*SonosTrackDescriptor `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// SonosListenEntryStream was auto-generated from WSDL.
type SonosListenEntryStream struct {
	XMLName        xml.Name `xml:"http://Mobile.ws.domoboxbusiness.com/ SonosListenEntryStream" json:"-" yaml:"-"`
	Session_key    string   `xml:"session_key,omitempty" json:"session_key,omitempty" yaml:"session_key,omitempty"`
	Device_key     string   `xml:"device_key,omitempty" json:"device_key,omitempty" yaml:"device_key,omitempty"`
	Device_key_src string   `xml:"device_key_src,omitempty" json:"device_key_src,omitempty" yaml:"device_key_src,omitempty"`
}

// SonosMakeNewGroup was auto-generated from WSDL.
type SonosMakeNewGroup struct {
	XMLName      xml.Name `xml:"http://Mobile.ws.domoboxbusiness.com/ SonosMakeNewGroup" json:"-" yaml:"-"`
	Session_key  string   `xml:"session_key,omitempty" json:"session_key,omitempty" yaml:"session_key,omitempty"`
	Devices_keys string   `xml:"devices_keys,omitempty" json:"devices_keys,omitempty" yaml:"devices_keys,omitempty"`
}

// SonosRemoveTrackFromPlaylist was auto-generated from WSDL.
type SonosRemoveTrackFromPlaylist struct {
	XMLName     xml.Name `xml:"http://Mobile.ws.domoboxbusiness.com/ SonosRemoveTrackFromPlaylist" json:"-" yaml:"-"`
	Session_key string   `xml:"session_key,omitempty" json:"session_key,omitempty" yaml:"session_key,omitempty"`
	Device_key  string   `xml:"device_key,omitempty" json:"device_key,omitempty" yaml:"device_key,omitempty"`
	Value       int      `xml:"value,omitempty" json:"value,omitempty" yaml:"value,omitempty"`
}

// SonosReorderTrack was auto-generated from WSDL.
type SonosReorderTrack struct {
	XMLName     xml.Name `xml:"http://Mobile.ws.domoboxbusiness.com/ SonosReorderTrack" json:"-" yaml:"-"`
	Session_key string   `xml:"session_key,omitempty" json:"session_key,omitempty" yaml:"session_key,omitempty"`
	Device_key  string   `xml:"device_key,omitempty" json:"device_key,omitempty" yaml:"device_key,omitempty"`
	Src         int      `xml:"src,omitempty" json:"src,omitempty" yaml:"src,omitempty"`
	Dest        int      `xml:"dest,omitempty" json:"dest,omitempty" yaml:"dest,omitempty"`
}

// SonosSetCurrentTrackNumber was auto-generated from WSDL.
type SonosSetCurrentTrackNumber struct {
	XMLName     xml.Name `xml:"http://Mobile.ws.domoboxbusiness.com/ SonosSetCurrentTrackNumber" json:"-" yaml:"-"`
	Session_key string   `xml:"session_key,omitempty" json:"session_key,omitempty" yaml:"session_key,omitempty"`
	Device_key  string   `xml:"device_key,omitempty" json:"device_key,omitempty" yaml:"device_key,omitempty"`
	Value       string   `xml:"value,omitempty" json:"value,omitempty" yaml:"value,omitempty"`
}

// SonosSetRandom was auto-generated from WSDL.
type SonosSetRandom struct {
	XMLName     xml.Name `xml:"http://Mobile.ws.domoboxbusiness.com/ SonosSetRandom" json:"-" yaml:"-"`
	Session_key string   `xml:"session_key,omitempty" json:"session_key,omitempty" yaml:"session_key,omitempty"`
	Device_key  string   `xml:"device_key,omitempty" json:"device_key,omitempty" yaml:"device_key,omitempty"`
	Value       bool     `xml:"value,omitempty" json:"value,omitempty" yaml:"value,omitempty"`
}

// SonosSetRepeat was auto-generated from WSDL.
type SonosSetRepeat struct {
	XMLName     xml.Name `xml:"http://Mobile.ws.domoboxbusiness.com/ SonosSetRepeat" json:"-" yaml:"-"`
	Session_key string   `xml:"session_key,omitempty" json:"session_key,omitempty" yaml:"session_key,omitempty"`
	Device_key  string   `xml:"device_key,omitempty" json:"device_key,omitempty" yaml:"device_key,omitempty"`
	Value       bool     `xml:"value,omitempty" json:"value,omitempty" yaml:"value,omitempty"`
}

// WSException was auto-generated from WSDL.
type WSException struct {
	ErrorCode             wsErrorCode `xml:"errorCode,omitempty" json:"errorCode,omitempty" yaml:"errorCode,omitempty"`
	ErrorTechnicalMessage string      `xml:"errorTechnicalMessage,omitempty" json:"errorTechnicalMessage,omitempty" yaml:"errorTechnicalMessage,omitempty"`
	ErrorMessage          string      `xml:"errorMessage,omitempty" json:"errorMessage,omitempty" yaml:"errorMessage,omitempty"`
}

// Action was auto-generated from WSDL.
type Action struct {
	Prop_clsid   string `xml:"prop_clsid,omitempty" json:"prop_clsid,omitempty" yaml:"prop_clsid,omitempty"`
	Action_clsid string `xml:"action_clsid,omitempty" json:"action_clsid,omitempty" yaml:"action_clsid,omitempty"`
	Descriptor   string `xml:"descriptor,omitempty" json:"descriptor,omitempty" yaml:"descriptor,omitempty"`
}

// Actions was auto-generated from WSDL.
type Actions struct {
	Action []*Action `xml:"action,omitempty" json:"action,omitempty" yaml:"action,omitempty"`
}

// AlertLog was auto-generated from WSDL.
type AlertLog struct {
	Label      string `xml:"label,omitempty" json:"label,omitempty" yaml:"label,omitempty"`
	Message    string `xml:"message,omitempty" json:"message,omitempty" yaml:"message,omitempty"`
	Date       string `xml:"date,omitempty" json:"date,omitempty" yaml:"date,omitempty"`
	Devc_clsid string `xml:"devc_clsid,omitempty" json:"devc_clsid,omitempty" yaml:"devc_clsid,omitempty"`
}

// AlertLogs was auto-generated from WSDL.
type AlertLogs struct {
	Session_key string      `xml:"session_key,omitempty" json:"session_key,omitempty" yaml:"session_key,omitempty"`
	Alert_log   []*AlertLog `xml:"alert_log,omitempty" json:"alert_log,omitempty" yaml:"alert_log,omitempty"`
}

// Bookmarks was auto-generated from WSDL.
type Bookmarks struct {
	Session_key string     `xml:"session_key,omitempty" json:"session_key,omitempty" yaml:"session_key,omitempty"`
	Devices     *Devices   `xml:"devices,omitempty" json:"devices,omitempty" yaml:"devices,omitempty"`
	Groups      *Groups    `xml:"groups,omitempty" json:"groups,omitempty" yaml:"groups,omitempty"`
	Scenarios   *Scenarios `xml:"scenarios,omitempty" json:"scenarios,omitempty" yaml:"scenarios,omitempty"`
}

// Camera was auto-generated from WSDL.
type Camera struct {
	Camera_key    string `xml:"camera_key,omitempty" json:"camera_key,omitempty" yaml:"camera_key,omitempty"`
	Camera_clsid  string `xml:"camera_clsid,omitempty" json:"camera_clsid,omitempty" yaml:"camera_clsid,omitempty"`
	Format_clsid  string `xml:"format_clsid,omitempty" json:"format_clsid,omitempty" yaml:"format_clsid,omitempty"`
	Label         string `xml:"label,omitempty" json:"label,omitempty" yaml:"label,omitempty"`
	Picture_key   string `xml:"picture_key,omitempty" json:"picture_key,omitempty" yaml:"picture_key,omitempty"`
	Url           string `xml:"url,omitempty" json:"url,omitempty" yaml:"url,omitempty"`
	Url_ext       string `xml:"url_ext,omitempty" json:"url_ext,omitempty" yaml:"url_ext,omitempty"`
	Login         string `xml:"login,omitempty" json:"login,omitempty" yaml:"login,omitempty"`
	Password      string `xml:"password,omitempty" json:"password,omitempty" yaml:"password,omitempty"`
	Url_left      string `xml:"url_left,omitempty" json:"url_left,omitempty" yaml:"url_left,omitempty"`
	Url_right     string `xml:"url_right,omitempty" json:"url_right,omitempty" yaml:"url_right,omitempty"`
	Url_up        string `xml:"url_up,omitempty" json:"url_up,omitempty" yaml:"url_up,omitempty"`
	Url_down      string `xml:"url_down,omitempty" json:"url_down,omitempty" yaml:"url_down,omitempty"`
	Url_preset1   string `xml:"url_preset1,omitempty" json:"url_preset1,omitempty" yaml:"url_preset1,omitempty"`
	Url_preset2   string `xml:"url_preset2,omitempty" json:"url_preset2,omitempty" yaml:"url_preset2,omitempty"`
	Url_preset3   string `xml:"url_preset3,omitempty" json:"url_preset3,omitempty" yaml:"url_preset3,omitempty"`
	Url_preset4   string `xml:"url_preset4,omitempty" json:"url_preset4,omitempty" yaml:"url_preset4,omitempty"`
	Url_zoom_up   string `xml:"url_zoom_up,omitempty" json:"url_zoom_up,omitempty" yaml:"url_zoom_up,omitempty"`
	Url_zoom_down string `xml:"url_zoom_down,omitempty" json:"url_zoom_down,omitempty" yaml:"url_zoom_down,omitempty"`
}

// Cameras was auto-generated from WSDL.
type Cameras struct {
	Session_key string    `xml:"session_key,omitempty" json:"session_key,omitempty" yaml:"session_key,omitempty"`
	Cameras     []*Camera `xml:"cameras,omitempty" json:"cameras,omitempty" yaml:"cameras,omitempty"`
}

// Categories was auto-generated from WSDL.
type Categories struct {
	Category []*Category `xml:"category,omitempty" json:"category,omitempty" yaml:"category,omitempty"`
}

// Category was auto-generated from WSDL.
type Category struct {
	Category_clsid string `xml:"category_clsid,omitempty" json:"category_clsid,omitempty" yaml:"category_clsid,omitempty"`
	Label          string `xml:"label,omitempty" json:"label,omitempty" yaml:"label,omitempty"`
	Picture_key    string `xml:"picture_key,omitempty" json:"picture_key,omitempty" yaml:"picture_key,omitempty"`
	Devices_count  string `xml:"devices_count,omitempty" json:"devices_count,omitempty" yaml:"devices_count,omitempty"`
}

// Consumptions was auto-generated from WSDL.
type Consumptions struct {
	Session_key  string       `xml:"session_key,omitempty" json:"session_key,omitempty" yaml:"session_key,omitempty"`
	Total_cons   *TotalCons   `xml:"total_cons,omitempty" json:"total_cons,omitempty" yaml:"total_cons,omitempty"`
	Rooms_cons   *RoomsCons   `xml:"rooms_cons,omitempty" json:"rooms_cons,omitempty" yaml:"rooms_cons,omitempty"`
	Devices_cons *DevicesCons `xml:"devices_cons,omitempty" json:"devices_cons,omitempty" yaml:"devices_cons,omitempty"`
}

// Device was auto-generated from WSDL.
type Device struct {
	Device_key     string   `xml:"device_key,omitempty" json:"device_key,omitempty" yaml:"device_key,omitempty"`
	Device_clsid   string   `xml:"device_clsid,omitempty" json:"device_clsid,omitempty" yaml:"device_clsid,omitempty"`
	Room_label     string   `xml:"room_label,omitempty" json:"room_label,omitempty" yaml:"room_label,omitempty"`
	Category_clsid string   `xml:"category_clsid,omitempty" json:"category_clsid,omitempty" yaml:"category_clsid,omitempty"`
	Label          string   `xml:"label,omitempty" json:"label,omitempty" yaml:"label,omitempty"`
	Resume         string   `xml:"resume,omitempty" json:"resume,omitempty" yaml:"resume,omitempty"`
	Alarm          bool     `xml:"alarm,omitempty" json:"alarm,omitempty" yaml:"alarm,omitempty"`
	Feedback_enbl  bool     `xml:"feedback_enbl,omitempty" json:"feedback_enbl,omitempty" yaml:"feedback_enbl,omitempty"`
	States         *States  `xml:"states,omitempty" json:"states,omitempty" yaml:"states,omitempty"`
	Actions        *Actions `xml:"actions,omitempty" json:"actions,omitempty" yaml:"actions,omitempty"`
}

// DeviceCons was auto-generated from WSDL.
type DeviceCons struct {
	Device_key     string  `xml:"device_key,omitempty" json:"device_key,omitempty" yaml:"device_key,omitempty"`
	Device_clsid   string  `xml:"device_clsid,omitempty" json:"device_clsid,omitempty" yaml:"device_clsid,omitempty"`
	Room_key       string  `xml:"room_key,omitempty" json:"room_key,omitempty" yaml:"room_key,omitempty"`
	Room_label     string  `xml:"room_label,omitempty" json:"room_label,omitempty" yaml:"room_label,omitempty"`
	Category_clsid string  `xml:"category_clsid,omitempty" json:"category_clsid,omitempty" yaml:"category_clsid,omitempty"`
	Label          string  `xml:"label,omitempty" json:"label,omitempty" yaml:"label,omitempty"`
	Picture_key    string  `xml:"picture_key,omitempty" json:"picture_key,omitempty" yaml:"picture_key,omitempty"`
	Value          float64 `xml:"value,omitempty" json:"value,omitempty" yaml:"value,omitempty"`
}

// Devices was auto-generated from WSDL.
type Devices struct {
	Session_key string    `xml:"session_key,omitempty" json:"session_key,omitempty" yaml:"session_key,omitempty"`
	Device      []*Device `xml:"device,omitempty" json:"device,omitempty" yaml:"device,omitempty"`
}

// DevicesCons was auto-generated from WSDL.
type DevicesCons struct {
	Device_cons []*DeviceCons `xml:"device_cons,omitempty" json:"device_cons,omitempty" yaml:"device_cons,omitempty"`
}

// DevicesGroup was auto-generated from WSDL.
type DevicesGroup struct {
	Device_label string  `xml:"device_label,omitempty" json:"device_label,omitempty" yaml:"device_label,omitempty"`
	Room_label   string  `xml:"room_label,omitempty" json:"room_label,omitempty" yaml:"room_label,omitempty"`
	Devc_clsid   string  `xml:"devc_clsid,omitempty" json:"devc_clsid,omitempty" yaml:"devc_clsid,omitempty"`
	Device_key   string  `xml:"device_key,omitempty" json:"device_key,omitempty" yaml:"device_key,omitempty"`
	States       *States `xml:"states,omitempty" json:"states,omitempty" yaml:"states,omitempty"`
}

// DevicesGroups was auto-generated from WSDL.
type DevicesGroups struct {
	Devices_group []*DevicesGroup `xml:"devices_group,omitempty" json:"devices_group,omitempty" yaml:"devices_group,omitempty"`
}

// Group was auto-generated from WSDL.
type Group struct {
	Group_key      string         `xml:"group_key,omitempty" json:"group_key,omitempty" yaml:"group_key,omitempty"`
	Label          string         `xml:"label,omitempty" json:"label,omitempty" yaml:"label,omitempty"`
	Resume         string         `xml:"resume,omitempty" json:"resume,omitempty" yaml:"resume,omitempty"`
	Family         string         `xml:"family,omitempty" json:"family,omitempty" yaml:"family,omitempty"`
	Actions        *Actions       `xml:"actions,omitempty" json:"actions,omitempty" yaml:"actions,omitempty"`
	States         *States        `xml:"states,omitempty" json:"states,omitempty" yaml:"states,omitempty"`
	Devices_groups *DevicesGroups `xml:"devices_groups,omitempty" json:"devices_groups,omitempty" yaml:"devices_groups,omitempty"`
}

// Groups was auto-generated from WSDL.
type Groups struct {
	Session_key string   `xml:"session_key,omitempty" json:"session_key,omitempty" yaml:"session_key,omitempty"`
	Group       []*Group `xml:"group,omitempty" json:"group,omitempty" yaml:"group,omitempty"`
}

// HistogramValue was auto-generated from WSDL.
type HistogramValue struct {
	Date   *LdDate   `xml:"date,omitempty" json:"date,omitempty" yaml:"date,omitempty"`
	Values []float64 `xml:"values,omitempty" json:"values,omitempty" yaml:"values,omitempty"`
}

// Infos was auto-generated from WSDL.
type Infos struct {
	Session_key string     `xml:"session_key,omitempty" json:"session_key,omitempty" yaml:"session_key,omitempty"`
	Rooms       *Rooms     `xml:"rooms,omitempty" json:"rooms,omitempty" yaml:"rooms,omitempty"`
	Bookmarks   *Bookmarks `xml:"bookmarks,omitempty" json:"bookmarks,omitempty" yaml:"bookmarks,omitempty"`
	Scenarios   *Scenarios `xml:"scenarios,omitempty" json:"scenarios,omitempty" yaml:"scenarios,omitempty"`
}

// InstantValue was auto-generated from WSDL.
type InstantValue struct {
	MaxScale float64 `xml:"maxScale,omitempty" json:"maxScale,omitempty" yaml:"maxScale,omitempty"`
	MaxValue float64 `xml:"maxValue,omitempty" json:"maxValue,omitempty" yaml:"maxValue,omitempty"`
	MinScale float64 `xml:"minScale,omitempty" json:"minScale,omitempty" yaml:"minScale,omitempty"`
	MinValue float64 `xml:"minValue,omitempty" json:"minValue,omitempty" yaml:"minValue,omitempty"`
	Value    float64 `xml:"value,omitempty" json:"value,omitempty" yaml:"value,omitempty"`
}

// LdDate was auto-generated from WSDL.
type LdDate struct {
	DayInMonth int   `xml:"dayInMonth,omitempty" json:"dayInMonth,omitempty" yaml:"dayInMonth,omitempty"`
	Month      month `xml:"month,omitempty" json:"month,omitempty" yaml:"month,omitempty"`
	Year       int   `xml:"year,omitempty" json:"year,omitempty" yaml:"year,omitempty"`
}

// MobileTranslationDescriptor was auto-generated from WSDL.
type MobileTranslationDescriptor struct {
	Base      string `xml:"base,omitempty" json:"base,omitempty" yaml:"base,omitempty"`
	Clsid     string `xml:"clsid,omitempty" json:"clsid,omitempty" yaml:"clsid,omitempty"`
	Clsid_lbl string `xml:"clsid_lbl,omitempty" json:"clsid_lbl,omitempty" yaml:"clsid_lbl,omitempty"`
	Label     string `xml:"label,omitempty" json:"label,omitempty" yaml:"label,omitempty"`
	Lang      string `xml:"lang,omitempty" json:"lang,omitempty" yaml:"lang,omitempty"`
}

// PieChartValue was auto-generated from WSDL.
type PieChartValue struct {
	Cold               float64 `xml:"cold,omitempty" json:"cold,omitempty" yaml:"cold,omitempty"`
	Cooling            float64 `xml:"cooling,omitempty" json:"cooling,omitempty" yaml:"cooling,omitempty"`
	Date               *LdDate `xml:"date,omitempty" json:"date,omitempty" yaml:"date,omitempty"`
	Domestic_hot_water float64 `xml:"domestic_hot_water,omitempty" json:"domestic_hot_water,omitempty" yaml:"domestic_hot_water,omitempty"`
	Heating            float64 `xml:"heating,omitempty" json:"heating,omitempty" yaml:"heating,omitempty"`
	Hot                float64 `xml:"hot,omitempty" json:"hot,omitempty" yaml:"hot,omitempty"`
	Plug               float64 `xml:"plug,omitempty" json:"plug,omitempty" yaml:"plug,omitempty"`
	Undetermined       float64 `xml:"undetermined,omitempty" json:"undetermined,omitempty" yaml:"undetermined,omitempty"`
}

// Room was auto-generated from WSDL.
type Room struct {
	Room_key    string `xml:"room_key,omitempty" json:"room_key,omitempty" yaml:"room_key,omitempty"`
	Label       string `xml:"label,omitempty" json:"label,omitempty" yaml:"label,omitempty"`
	Picture_key string `xml:"picture_key,omitempty" json:"picture_key,omitempty" yaml:"picture_key,omitempty"`
}

// RoomCons was auto-generated from WSDL.
type RoomCons struct {
	Room_key    string  `xml:"room_key,omitempty" json:"room_key,omitempty" yaml:"room_key,omitempty"`
	Room_label  string  `xml:"room_label,omitempty" json:"room_label,omitempty" yaml:"room_label,omitempty"`
	Picture_key string  `xml:"picture_key,omitempty" json:"picture_key,omitempty" yaml:"picture_key,omitempty"`
	Value       float64 `xml:"value,omitempty" json:"value,omitempty" yaml:"value,omitempty"`
}

// Rooms was auto-generated from WSDL.
type Rooms struct {
	Room []*Room `xml:"room,omitempty" json:"room,omitempty" yaml:"room,omitempty"`
}

// RoomsCons was auto-generated from WSDL.
type RoomsCons struct {
	Room_cons []*RoomCons `xml:"room_cons,omitempty" json:"room_cons,omitempty" yaml:"room_cons,omitempty"`
}

// RtPageDescriptor was auto-generated from WSDL.
type RtPageDescriptor struct {
	Chart_types   []chartType   `xml:"chart_types,omitempty" json:"chart_types,omitempty" yaml:"chart_types,omitempty"`
	Instant_types []instantType `xml:"instant_types,omitempty" json:"instant_types,omitempty" yaml:"instant_types,omitempty"`
	Types         []rtDataType  `xml:"types,omitempty" json:"types,omitempty" yaml:"types,omitempty"`
}

// Scenario was auto-generated from WSDL.
type Scenario struct {
	Scenario_key      string `xml:"scenario_key,omitempty" json:"scenario_key,omitempty" yaml:"scenario_key,omitempty"`
	Label             string `xml:"label,omitempty" json:"label,omitempty" yaml:"label,omitempty"`
	Resume            string `xml:"resume,omitempty" json:"resume,omitempty" yaml:"resume,omitempty"`
	Next_exe_date     string `xml:"next_exe_date,omitempty" json:"next_exe_date,omitempty" yaml:"next_exe_date,omitempty"`
	Activate_state    string `xml:"activate_state,omitempty" json:"activate_state,omitempty" yaml:"activate_state,omitempty"`
	Authorization_act string `xml:"authorization_act,omitempty" json:"authorization_act,omitempty" yaml:"authorization_act,omitempty"`
	Authorization_des string `xml:"authorization_des,omitempty" json:"authorization_des,omitempty" yaml:"authorization_des,omitempty"`
	Last_exec         string `xml:"last_exec,omitempty" json:"last_exec,omitempty" yaml:"last_exec,omitempty"`
	Tasks             *Tasks `xml:"tasks,omitempty" json:"tasks,omitempty" yaml:"tasks,omitempty"`
}

// Scenarios was auto-generated from WSDL.
type Scenarios struct {
	Session_key string      `xml:"session_key,omitempty" json:"session_key,omitempty" yaml:"session_key,omitempty"`
	Scenario    []*Scenario `xml:"scenario,omitempty" json:"scenario,omitempty" yaml:"scenario,omitempty"`
}

// Site was auto-generated from WSDL.
type Site struct {
	Site_key    string `xml:"site_key,omitempty" json:"site_key,omitempty" yaml:"site_key,omitempty"`
	Label       string `xml:"label,omitempty" json:"label,omitempty" yaml:"label,omitempty"`
	Picture_key string `xml:"picture_key,omitempty" json:"picture_key,omitempty" yaml:"picture_key,omitempty"`
}

// Sites was auto-generated from WSDL.
type Sites struct {
	Site []*Site `xml:"site,omitempty" json:"site,omitempty" yaml:"site,omitempty"`
}

// SonosGroupDescriptor was auto-generated from WSDL.
type SonosGroupDescriptor struct {
	Coordinator_key string `xml:"coordinator_key,omitempty" json:"coordinator_key,omitempty" yaml:"coordinator_key,omitempty"`
	Device_key      string `xml:"device_key,omitempty" json:"device_key,omitempty" yaml:"device_key,omitempty"`
	Sonos_label     string `xml:"sonos_label,omitempty" json:"sonos_label,omitempty" yaml:"sonos_label,omitempty"`
}

// SonosMobileGroupDescriptor was auto-generated from WSDL.
type SonosMobileGroupDescriptor struct {
	Album           string `xml:"album,omitempty" json:"album,omitempty" yaml:"album,omitempty"`
	Artist          string `xml:"artist,omitempty" json:"artist,omitempty" yaml:"artist,omitempty"`
	Coordinator_key string `xml:"coordinator_key,omitempty" json:"coordinator_key,omitempty" yaml:"coordinator_key,omitempty"`
	Cover_link      string `xml:"cover_link,omitempty" json:"cover_link,omitempty" yaml:"cover_link,omitempty"`
	Device_key      string `xml:"device_key,omitempty" json:"device_key,omitempty" yaml:"device_key,omitempty"`
	Is_playing      bool   `xml:"is_playing,omitempty" json:"is_playing,omitempty" yaml:"is_playing,omitempty"`
	Sonos_label     string `xml:"sonos_label,omitempty" json:"sonos_label,omitempty" yaml:"sonos_label,omitempty"`
	Title           string `xml:"title,omitempty" json:"title,omitempty" yaml:"title,omitempty"`
}

// SonosTrackDescriptor was auto-generated from WSDL.
type SonosTrackDescriptor struct {
	Album       string `xml:"album,omitempty" json:"album,omitempty" yaml:"album,omitempty"`
	Artist      string `xml:"artist,omitempty" json:"artist,omitempty" yaml:"artist,omitempty"`
	Container   bool   `xml:"container,omitempty" json:"container,omitempty" yaml:"container,omitempty"`
	Cover       string `xml:"cover,omitempty" json:"cover,omitempty" yaml:"cover,omitempty"`
	Description string `xml:"description,omitempty" json:"description,omitempty" yaml:"description,omitempty"`
	Device_key  string `xml:"device_key,omitempty" json:"device_key,omitempty" yaml:"device_key,omitempty"`
	Id          string `xml:"id,omitempty" json:"id,omitempty" yaml:"id,omitempty"`
	Metadata    string `xml:"metadata,omitempty" json:"metadata,omitempty" yaml:"metadata,omitempty"`
	Number      int    `xml:"number,omitempty" json:"number,omitempty" yaml:"number,omitempty"`
	Rincon_uri  string `xml:"rincon_uri,omitempty" json:"rincon_uri,omitempty" yaml:"rincon_uri,omitempty"`
	Title       string `xml:"title,omitempty" json:"title,omitempty" yaml:"title,omitempty"`
	Track_type  int    `xml:"track_type,omitempty" json:"track_type,omitempty" yaml:"track_type,omitempty"`
	Type_str    string `xml:"type_str,omitempty" json:"type_str,omitempty" yaml:"type_str,omitempty"`
}

// State was auto-generated from WSDL.
type State struct {
	State_clsid string  `xml:"state_clsid,omitempty" json:"state_clsid,omitempty" yaml:"state_clsid,omitempty"`
	Type        string  `xml:"type,omitempty" json:"type,omitempty" yaml:"type,omitempty"`
	Label       string  `xml:"label,omitempty" json:"label,omitempty" yaml:"label,omitempty"`
	Values      *Values `xml:"values,omitempty" json:"values,omitempty" yaml:"values,omitempty"`
}

// States was auto-generated from WSDL.
type States struct {
	State []*State `xml:"state,omitempty" json:"state,omitempty" yaml:"state,omitempty"`
}

// Task was auto-generated from WSDL.
type Task struct {
	Target     string `xml:"target,omitempty" json:"target,omitempty" yaml:"target,omitempty"`
	Action     string `xml:"action,omitempty" json:"action,omitempty" yaml:"action,omitempty"`
	Clsid      string `xml:"clsid,omitempty" json:"clsid,omitempty" yaml:"clsid,omitempty"`
	Room_label string `xml:"room_label,omitempty" json:"room_label,omitempty" yaml:"room_label,omitempty"`
	Offset     int    `xml:"offset,omitempty" json:"offset,omitempty" yaml:"offset,omitempty"`
}

// Tasks was auto-generated from WSDL.
type Tasks struct {
	Task []*Task `xml:"task,omitempty" json:"task,omitempty" yaml:"task,omitempty"`
}

// TotalCons was auto-generated from WSDL.
type TotalCons struct {
	Value    float64 `xml:"value,omitempty" json:"value,omitempty" yaml:"value,omitempty"`
	Prix     string  `xml:"prix,omitempty" json:"prix,omitempty" yaml:"prix,omitempty"`
	Currency string  `xml:"currency,omitempty" json:"currency,omitempty" yaml:"currency,omitempty"`
}

// TotalValue was auto-generated from WSDL.
type TotalValue struct {
	Date        *LdDate `xml:"date,omitempty" json:"date,omitempty" yaml:"date,omitempty"`
	Date_reset  *LdDate `xml:"date_reset,omitempty" json:"date_reset,omitempty" yaml:"date_reset,omitempty"`
	Value       float64 `xml:"value,omitempty" json:"value,omitempty" yaml:"value,omitempty"`
	Value_reset float64 `xml:"value_reset,omitempty" json:"value_reset,omitempty" yaml:"value_reset,omitempty"`
}

// User was auto-generated from WSDL.
type User struct {
	User_key    string `xml:"user_key,omitempty" json:"user_key,omitempty" yaml:"user_key,omitempty"`
	Nickname    string `xml:"nickname,omitempty" json:"nickname,omitempty" yaml:"nickname,omitempty"`
	Picture_key string `xml:"picture_key,omitempty" json:"picture_key,omitempty" yaml:"picture_key,omitempty"`
}

// Users was auto-generated from WSDL.
type Users struct {
	User []*User `xml:"user,omitempty" json:"user,omitempty" yaml:"user,omitempty"`
}

// Value was auto-generated from WSDL.
type Value struct {
	Index       string `xml:"index,omitempty" json:"index,omitempty" yaml:"index,omitempty"`
	Value       string `xml:"value,omitempty" json:"value,omitempty" yaml:"value,omitempty"`
	Description string `xml:"description,omitempty" json:"description,omitempty" yaml:"description,omitempty"`
	Unit        string `xml:"unit,omitempty" json:"unit,omitempty" yaml:"unit,omitempty"`
	Label       string `xml:"label,omitempty" json:"label,omitempty" yaml:"label,omitempty"`
}

// Values was auto-generated from WSDL.
type Values struct {
	Value []*Value `xml:"value,omitempty" json:"value,omitempty" yaml:"value,omitempty"`
}

// VarunaInfos was auto-generated from WSDL.
type VarunaInfos struct {
	Date    *ZonedDate `xml:"date,omitempty" json:"date,omitempty" yaml:"date,omitempty"`
	Message string     `xml:"message,omitempty" json:"message,omitempty" yaml:"message,omitempty"`
}

// Voicedevice was auto-generated from WSDL.
type Voicedevice struct {
	Name       string               `xml:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
	Device_key string               `xml:"device_key,omitempty" json:"device_key,omitempty" yaml:"device_key,omitempty"`
	Actions    []*Voicedeviceaction `xml:"actions,omitempty" json:"actions,omitempty" yaml:"actions,omitempty"`
}

// Voicedeviceaction was auto-generated from WSDL.
type Voicedeviceaction struct {
	Prop_clsid   string `xml:"prop_clsid,omitempty" json:"prop_clsid,omitempty" yaml:"prop_clsid,omitempty"`
	Action_clsid string `xml:"action_clsid,omitempty" json:"action_clsid,omitempty" yaml:"action_clsid,omitempty"`
}

// Voiceroom was auto-generated from WSDL.
type Voiceroom struct {
	Name    string         `xml:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
	Devices []*Voicedevice `xml:"devices,omitempty" json:"devices,omitempty" yaml:"devices,omitempty"`
}

// Voicescenario was auto-generated from WSDL.
type Voicescenario struct {
	Name         string `xml:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
	Scenario_key string `xml:"scenario_key,omitempty" json:"scenario_key,omitempty" yaml:"scenario_key,omitempty"`
}

// Voicestruct was auto-generated from WSDL.
type Voicestruct struct {
	Rooms     []*Voiceroom     `xml:"rooms,omitempty" json:"rooms,omitempty" yaml:"rooms,omitempty"`
	Scenarios []*Voicescenario `xml:"scenarios,omitempty" json:"scenarios,omitempty" yaml:"scenarios,omitempty"`
}

// XxdAlarmFaultDescriptor was auto-generated from WSDL.
type XxdAlarmFaultDescriptor struct {
	Index       int    `xml:"index,omitempty" json:"index,omitempty" yaml:"index,omitempty"`
	Label       string `xml:"label,omitempty" json:"label,omitempty" yaml:"label,omitempty"`
	ProductId   string `xml:"productId,omitempty" json:"productId,omitempty" yaml:"productId,omitempty"`
	ProductType string `xml:"productType,omitempty" json:"productType,omitempty" yaml:"productType,omitempty"`
	Zone        string `xml:"zone,omitempty" json:"zone,omitempty" yaml:"zone,omitempty"`
}

// XxdAlarmFaultsStatisticsDescriptor was auto-generated from WSDL.
type XxdAlarmFaultsStatisticsDescriptor struct {
	Inhibitions          int `xml:"inhibitions,omitempty" json:"inhibitions,omitempty" yaml:"inhibitions,omitempty"`
	OpenEntrance         int `xml:"openEntrance,omitempty" json:"openEntrance,omitempty" yaml:"openEntrance,omitempty"`
	Others               int `xml:"others,omitempty" json:"others,omitempty" yaml:"others,omitempty"`
	SelfProtection       int `xml:"selfProtection,omitempty" json:"selfProtection,omitempty" yaml:"selfProtection,omitempty"`
	UnacknowledgedEvents int `xml:"unacknowledgedEvents,omitempty" json:"unacknowledgedEvents,omitempty" yaml:"unacknowledgedEvents,omitempty"`
}

// XxdAlarmHistory was auto-generated from WSDL.
type XxdAlarmHistory struct {
	Index          int      `xml:"index,omitempty" json:"index,omitempty" yaml:"index,omitempty"`
	Label          string   `xml:"label,omitempty" json:"label,omitempty" yaml:"label,omitempty"`
	OriginatorId   string   `xml:"originatorId,omitempty" json:"originatorId,omitempty" yaml:"originatorId,omitempty"`
	OriginatorType string   `xml:"originatorType,omitempty" json:"originatorType,omitempty" yaml:"originatorType,omitempty"`
	Timestamp      int64    `xml:"timestamp,omitempty" json:"timestamp,omitempty" yaml:"timestamp,omitempty"`
	Zones          []string `xml:"zones,omitempty" json:"zones,omitempty" yaml:"zones,omitempty"`
}

// ZonedDate was auto-generated from WSDL.
type ZonedDate struct {
	DayInMonth        int   `xml:"dayInMonth,omitempty" json:"dayInMonth,omitempty" yaml:"dayInMonth,omitempty"`
	GMTOffsetInMinute int   `xml:"GMTOffsetInMinute,omitempty" json:"GMTOffsetInMinute,omitempty" yaml:"GMTOffsetInMinute,omitempty"`
	Hour              int   `xml:"hour,omitempty" json:"hour,omitempty" yaml:"hour,omitempty"`
	MilliSecond       int   `xml:"milliSecond,omitempty" json:"milliSecond,omitempty" yaml:"milliSecond,omitempty"`
	Minute            int   `xml:"minute,omitempty" json:"minute,omitempty" yaml:"minute,omitempty"`
	Month             month `xml:"month,omitempty" json:"month,omitempty" yaml:"month,omitempty"`
	Second            int   `xml:"second,omitempty" json:"second,omitempty" yaml:"second,omitempty"`
	Year              int   `xml:"year,omitempty" json:"year,omitempty" yaml:"year,omitempty"`
}

// mobile implements the Mobile interface.
type mobile struct {
	cli *soap.Client
}

// RegisterNotifPhoneAndroid was auto-generated from WSDL.
func (p *mobile) RegisterNotifPhoneAndroid(α *RegisterNotifPhoneAndroid) (β *RegisterNotifPhoneAndroidResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M RegisterNotifPhoneAndroidResponse `xml:"RegisterNotifPhoneAndroidResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// RegisterNotifIpad2 was auto-generated from WSDL.
func (p *mobile) RegisterNotifIpad2(α *RegisterNotifIpad2) (β *RegisterNotifIpad2Response, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M RegisterNotifIpad2Response `xml:"RegisterNotifIpad2Response"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// SonosGetAllGroups was auto-generated from WSDL.
func (p *mobile) SonosGetAllGroups(α *SonosGetAllGroups) (β *SonosGetAllGroupsResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M SonosGetAllGroupsResponse `xml:"SonosGetAllGroupsResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// GetVarunaAlarms was auto-generated from WSDL.
func (p *mobile) GetVarunaAlarms(α *GetVarunaAlarms) (β *GetVarunaAlarmsResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M GetVarunaAlarmsResponse `xml:"GetVarunaAlarmsResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// SonosMakeNewGroup was auto-generated from WSDL.
func (p *mobile) SonosMakeNewGroup(α *SonosMakeNewGroup) (β *SonosMakeNewGroupResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M SonosMakeNewGroupResponse `xml:"SonosMakeNewGroupResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// GetDeviceInstantane was auto-generated from WSDL.
func (p *mobile) GetDeviceInstantane(α *GetDeviceInstantane) (β *GetDeviceInstantaneResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M GetDeviceInstantaneResponse `xml:"GetDeviceInstantaneResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// LoginBookmarks was auto-generated from WSDL.
func (p *mobile) LoginBookmarks(α *LoginBookmarks) (β *LoginBookmarksResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M LoginBookmarksResponse `xml:"LoginBookmarksResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// LoginConsumptions was auto-generated from WSDL.
func (p *mobile) LoginConsumptions(α *LoginConsumptions) (β *LoginConsumptionsResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M LoginConsumptionsResponse `xml:"LoginConsumptionsResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// LoginAlertLogs was auto-generated from WSDL.
func (p *mobile) LoginAlertLogs(α *LoginAlertLogs) (β *LoginAlertLogsResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M LoginAlertLogsResponse `xml:"LoginAlertLogsResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// GetInstantTypeList was auto-generated from WSDL.
func (p *mobile) GetInstantTypeList(α *GetInstantTypeList) (β *GetInstantTypeListResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M GetInstantTypeListResponse `xml:"GetInstantTypeListResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// KeepAlive was auto-generated from WSDL.
func (p *mobile) KeepAlive(α *KeepAlive) (β *KeepAliveResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M KeepAliveResponse `xml:"KeepAliveResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// GetVarunaEvents was auto-generated from WSDL.
func (p *mobile) GetVarunaEvents(α *GetVarunaEvents) (β *GetVarunaEventsResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M GetVarunaEventsResponse `xml:"GetVarunaEventsResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// GetGroup was auto-generated from WSDL.
func (p *mobile) GetGroup(α *GetGroup) (β *GetGroupResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M GetGroupResponse `xml:"GetGroupResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// SonosSetRandom was auto-generated from WSDL.
func (p *mobile) SonosSetRandom(α *SonosSetRandom) (β *SonosSetRandomResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M SonosSetRandomResponse `xml:"SonosSetRandomResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// SonosDeleteFavorite was auto-generated from WSDL.
func (p *mobile) SonosDeleteFavorite(α *SonosDeleteFavorite) (β *SonosDeleteFavoriteResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M SonosDeleteFavoriteResponse `xml:"SonosDeleteFavoriteResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// GetDeviceTranslations was auto-generated from WSDL.
func (p *mobile) GetDeviceTranslations(α *GetDeviceTranslations) (β *GetDeviceTranslationsResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M GetDeviceTranslationsResponse `xml:"GetDeviceTranslationsResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// ResetInstantValue was auto-generated from WSDL.
func (p *mobile) ResetInstantValue(α *ResetInstantValue) (β *ResetInstantValueResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M ResetInstantValueResponse `xml:"ResetInstantValueResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// GetDevices was auto-generated from WSDL.
func (p *mobile) GetDevices(α *GetDevices) (β *GetDevicesResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M GetDevicesResponse `xml:"GetDevicesResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// RegisterNotifTabletteAndroid was auto-generated from WSDL.
func (p *mobile) RegisterNotifTabletteAndroid(α *RegisterNotifTabletteAndroid) (β *RegisterNotifTabletteAndroidResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M RegisterNotifTabletteAndroidResponse `xml:"RegisterNotifTabletteAndroidResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// GetCameras was auto-generated from WSDL.
func (p *mobile) GetCameras(α *GetCameras) (β *GetCamerasResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M GetCamerasResponse `xml:"GetCamerasResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// SonosAddAndPlay was auto-generated from WSDL.
func (p *mobile) SonosAddAndPlay(α *SonosAddAndPlay) (β *SonosAddAndPlayResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M SonosAddAndPlayResponse `xml:"SonosAddAndPlayResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// GetAlertLogs was auto-generated from WSDL.
func (p *mobile) GetAlertLogs(α *GetAlertLogs) (β *GetAlertLogsResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M GetAlertLogsResponse `xml:"GetAlertLogsResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// GetAlarmFaultsStatistics was auto-generated from WSDL.
func (p *mobile) GetAlarmFaultsStatistics(α *GetAlarmFaultsStatistics) (β *GetAlarmFaultsStatisticsResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M GetAlarmFaultsStatisticsResponse `xml:"GetAlarmFaultsStatisticsResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// ExecuteAction was auto-generated from WSDL.
func (p *mobile) ExecuteAction(α *ExecuteAction) (β *ExecuteActionResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M ExecuteActionResponse `xml:"ExecuteActionResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// Login was auto-generated from WSDL.
func (p *mobile) Login(α *Login) (β *LoginResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M LoginResponse `xml:"LoginResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// SonosRemoveTrackFromPlaylist was auto-generated from WSDL.
func (p *mobile) SonosRemoveTrackFromPlaylist(α *SonosRemoveTrackFromPlaylist) (β *SonosRemoveTrackFromPlaylistResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M SonosRemoveTrackFromPlaylistResponse `xml:"SonosRemoveTrackFromPlaylistResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// LoginDevicesFromCatg was auto-generated from WSDL.
func (p *mobile) LoginDevicesFromCatg(α *LoginDevicesFromCatg) (β *LoginDevicesFromCatgResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M LoginDevicesFromCatgResponse `xml:"LoginDevicesFromCatgResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// ClearAndroidNotificationDataBase was auto-generated from WSDL.
func (p *mobile) ClearAndroidNotificationDataBase(α *ClearAndroidNotificationDataBase) (β *ClearAndroidNotificationDataBaseResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M ClearAndroidNotificationDataBaseResponse `xml:"ClearAndroidNotificationDataBaseResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// SonosExplore was auto-generated from WSDL.
func (p *mobile) SonosExplore(α *SonosExplore) (β *SonosExploreResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M SonosExploreResponse `xml:"SonosExploreResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// LoginDevices was auto-generated from WSDL.
func (p *mobile) LoginDevices(α *LoginDevices) (β *LoginDevicesResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M LoginDevicesResponse `xml:"LoginDevicesResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// GetHistValue was auto-generated from WSDL.
func (p *mobile) GetHistValue(α *GetHistValue) (β *GetHistValueResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M GetHistValueResponse `xml:"GetHistValueResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// GetTranslations was auto-generated from WSDL.
func (p *mobile) GetTranslations(α *GetTranslations) (β *GetTranslationsResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M GetTranslationsResponse `xml:"GetTranslationsResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// GetStateTranslations was auto-generated from WSDL.
func (p *mobile) GetStateTranslations(α *GetStateTranslations) (β *GetStateTranslationsResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M GetStateTranslationsResponse `xml:"GetStateTranslationsResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// SonosSetCurrentTrackNumber was auto-generated from WSDL.
func (p *mobile) SonosSetCurrentTrackNumber(α *SonosSetCurrentTrackNumber) (β *SonosSetCurrentTrackNumberResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M SonosSetCurrentTrackNumberResponse `xml:"SonosSetCurrentTrackNumberResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// SonosDeleteList was auto-generated from WSDL.
func (p *mobile) SonosDeleteList(α *SonosDeleteList) (β *SonosDeleteListResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M SonosDeleteListResponse `xml:"SonosDeleteListResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// GetScenarios was auto-generated from WSDL.
func (p *mobile) GetScenarios(α *GetScenarios) (β *GetScenariosResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M GetScenariosResponse `xml:"GetScenariosResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// LoginExecuteAction was auto-generated from WSDL.
func (p *mobile) LoginExecuteAction(α *LoginExecuteAction) (β *LoginExecuteActionResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M LoginExecuteActionResponse `xml:"LoginExecuteActionResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// GetDeviceState was auto-generated from WSDL.
func (p *mobile) GetDeviceState(α *GetDeviceState) (β *GetDeviceStateResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M GetDeviceStateResponse `xml:"GetDeviceStateResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// GetAlarmeAccessLevel was auto-generated from WSDL.
func (p *mobile) GetAlarmeAccessLevel(α *GetAlarmeAccessLevel) (β *GetAlarmeAccessLevelResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M GetAlarmeAccessLevelResponse `xml:"GetAlarmeAccessLevelResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// GetAlarmHistoryObjectList was auto-generated from WSDL.
func (p *mobile) GetAlarmHistoryObjectList(α *GetAlarmHistoryObjectList) (β *GetAlarmHistoryObjectListResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M GetAlarmHistoryObjectListResponse `xml:"GetAlarmHistoryObjectListResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// ResetNotif was auto-generated from WSDL.
func (p *mobile) ResetNotif(α *ResetNotif) (β *ResetNotifResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M ResetNotifResponse `xml:"ResetNotifResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// LoginGroup was auto-generated from WSDL.
func (p *mobile) LoginGroup(α *LoginGroup) (β *LoginGroupResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M LoginGroupResponse `xml:"LoginGroupResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// GetRTPageDescriptor was auto-generated from WSDL.
func (p *mobile) GetRTPageDescriptor(α *GetRTPageDescriptor) (β *GetRTPageDescriptorResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M GetRTPageDescriptorResponse `xml:"GetRTPageDescriptorResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// SonosGetPlaylist was auto-generated from WSDL.
func (p *mobile) SonosGetPlaylist(α *SonosGetPlaylist) (β *SonosGetPlaylistResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M SonosGetPlaylistResponse `xml:"SonosGetPlaylistResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// GetActionTranslations was auto-generated from WSDL.
func (p *mobile) GetActionTranslations(α *GetActionTranslations) (β *GetActionTranslationsResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M GetActionTranslationsResponse `xml:"GetActionTranslationsResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// RefreshVarunaEvents was auto-generated from WSDL.
func (p *mobile) RefreshVarunaEvents(α *RefreshVarunaEvents) (β *RefreshVarunaEventsResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M RefreshVarunaEventsResponse `xml:"RefreshVarunaEventsResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// RegisterNotif was auto-generated from WSDL.
func (p *mobile) RegisterNotif(α *RegisterNotif) (β *RegisterNotifResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M RegisterNotifResponse `xml:"RegisterNotifResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// GetRooms was auto-generated from WSDL.
func (p *mobile) GetRooms(α *GetRooms) (β *GetRoomsResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M GetRoomsResponse `xml:"GetRoomsResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// RegisterNotifIpad was auto-generated from WSDL.
func (p *mobile) RegisterNotifIpad(α *RegisterNotifIpad) (β *RegisterNotifIpadResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M RegisterNotifIpadResponse `xml:"RegisterNotifIpadResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// SonosListenEntryStream was auto-generated from WSDL.
func (p *mobile) SonosListenEntryStream(α *SonosListenEntryStream) (β *SonosListenEntryStreamResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M SonosListenEntryStreamResponse `xml:"SonosListenEntryStreamResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// GetVoiceObjects was auto-generated from WSDL.
func (p *mobile) GetVoiceObjects(α *GetVoiceObjects) (β *GetVoiceObjectsResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M GetVoiceObjectsResponse `xml:"GetVoiceObjectsResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// SonosSetRepeat was auto-generated from WSDL.
func (p *mobile) SonosSetRepeat(α *SonosSetRepeat) (β *SonosSetRepeatResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M SonosSetRepeatResponse `xml:"SonosSetRepeatResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// LoginDeviceState was auto-generated from WSDL.
func (p *mobile) LoginDeviceState(α *LoginDeviceState) (β *LoginDeviceStateResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M LoginDeviceStateResponse `xml:"LoginDeviceStateResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// LoginInfos was auto-generated from WSDL.
func (p *mobile) LoginInfos(α *LoginInfos) (β *LoginInfosResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M LoginInfosResponse `xml:"LoginInfosResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// GetPieChartValue was auto-generated from WSDL.
func (p *mobile) GetPieChartValue(α *GetPieChartValue) (β *GetPieChartValueResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M GetPieChartValueResponse `xml:"GetPieChartValueResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// GetCategories was auto-generated from WSDL.
func (p *mobile) GetCategories(α *GetCategories) (β *GetCategoriesResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M GetCategoriesResponse `xml:"GetCategoriesResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// CancelMobilePicture was auto-generated from WSDL.
func (p *mobile) CancelMobilePicture(α *CancelMobilePicture) (β *CancelMobilePictureResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M CancelMobilePictureResponse `xml:"CancelMobilePictureResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// LoginCameras was auto-generated from WSDL.
func (p *mobile) LoginCameras(α *LoginCameras) (β *LoginCamerasResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M LoginCamerasResponse `xml:"LoginCamerasResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// SonosAddFavorite was auto-generated from WSDL.
func (p *mobile) SonosAddFavorite(α *SonosAddFavorite) (β *SonosAddFavoriteResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M SonosAddFavoriteResponse `xml:"SonosAddFavoriteResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// GetGroups was auto-generated from WSDL.
func (p *mobile) GetGroups(α *GetGroups) (β *GetGroupsResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M GetGroupsResponse `xml:"GetGroupsResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// LoginScenarios was auto-generated from WSDL.
func (p *mobile) LoginScenarios(α *LoginScenarios) (β *LoginScenariosResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M LoginScenariosResponse `xml:"LoginScenariosResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// SonosReorderTrack was auto-generated from WSDL.
func (p *mobile) SonosReorderTrack(α *SonosReorderTrack) (β *SonosReorderTrackResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M SonosReorderTrackResponse `xml:"SonosReorderTrackResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// GetRTDataTypeList was auto-generated from WSDL.
func (p *mobile) GetRTDataTypeList(α *GetRTDataTypeList) (β *GetRTDataTypeListResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M GetRTDataTypeListResponse `xml:"GetRTDataTypeListResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// ApplyMobilePicture was auto-generated from WSDL.
func (p *mobile) ApplyMobilePicture(α *ApplyMobilePicture) (β *ApplyMobilePictureResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M ApplyMobilePictureResponse `xml:"ApplyMobilePictureResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// SonosGetLineInList was auto-generated from WSDL.
func (p *mobile) SonosGetLineInList(α *SonosGetLineInList) (β *SonosGetLineInListResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M SonosGetLineInListResponse `xml:"SonosGetLineInListResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// LoginGroups was auto-generated from WSDL.
func (p *mobile) LoginGroups(α *LoginGroups) (β *LoginGroupsResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M LoginGroupsResponse `xml:"LoginGroupsResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// GetTotalDataValue was auto-generated from WSDL.
func (p *mobile) GetTotalDataValue(α *GetTotalDataValue) (β *GetTotalDataValueResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M GetTotalDataValueResponse `xml:"GetTotalDataValueResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// GetPicture was auto-generated from WSDL.
func (p *mobile) GetPicture(α *GetPicture) (β *GetPictureResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M GetPictureResponse `xml:"GetPictureResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// GetConsumptions was auto-generated from WSDL.
func (p *mobile) GetConsumptions(α *GetConsumptions) (β *GetConsumptionsResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M GetConsumptionsResponse `xml:"GetConsumptionsResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// NewLoginInfos was auto-generated from WSDL.
func (p *mobile) NewLoginInfos(α *NewLoginInfos) (β *NewLoginInfosResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M NewLoginInfosResponse `xml:"NewLoginInfosResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// GetUsers was auto-generated from WSDL.
func (p *mobile) GetUsers(α *GetUsers) (β *GetUsersResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M GetUsersResponse `xml:"GetUsersResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// GetTranslationsREST was auto-generated from WSDL.
func (p *mobile) GetTranslationsREST(α *GetTranslationsREST) (β *GetTranslationsRESTResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M GetTranslationsRESTResponse `xml:"GetTranslationsRESTResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// GetBookmarks was auto-generated from WSDL.
func (p *mobile) GetBookmarks(α *GetBookmarks) (β *GetBookmarksResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M GetBookmarksResponse `xml:"GetBookmarksResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// GetRoomsNew was auto-generated from WSDL.
func (p *mobile) GetRoomsNew(α *GetRoomsNew) (β *GetRoomsNewResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M GetRoomsNewResponse `xml:"GetRoomsNewResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// GetCategoriesTranslations was auto-generated from WSDL.
func (p *mobile) GetCategoriesTranslations(α *GetCategoriesTranslations) (β *GetCategoriesTranslationsResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M GetCategoriesTranslationsResponse `xml:"GetCategoriesTranslationsResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// GetDevicesFromCatg was auto-generated from WSDL.
func (p *mobile) GetDevicesFromCatg(α *GetDevicesFromCatg) (β *GetDevicesFromCatgResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M GetDevicesFromCatgResponse `xml:"GetDevicesFromCatgResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// GetRTDataCategoriesList was auto-generated from WSDL.
func (p *mobile) GetRTDataCategoriesList(α *GetRTDataCategoriesList) (β *GetRTDataCategoriesListResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M GetRTDataCategoriesListResponse `xml:"GetRTDataCategoriesListResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// ResetTotalValue was auto-generated from WSDL.
func (p *mobile) ResetTotalValue(α *ResetTotalValue) (β *ResetTotalValueResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M ResetTotalValueResponse `xml:"ResetTotalValueResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// GetSites was auto-generated from WSDL.
func (p *mobile) GetSites(α *GetSites) (β *GetSitesResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M GetSitesResponse `xml:"GetSitesResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// SonosAddTrackInPlaylist was auto-generated from WSDL.
func (p *mobile) SonosAddTrackInPlaylist(α *SonosAddTrackInPlaylist) (β *SonosAddTrackInPlaylistResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M SonosAddTrackInPlaylistResponse `xml:"SonosAddTrackInPlaylistResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// GetAlarmFaultObjectList was auto-generated from WSDL.
func (p *mobile) GetAlarmFaultObjectList(α *GetAlarmFaultObjectList) (β *GetAlarmFaultObjectListResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M GetAlarmFaultObjectListResponse `xml:"GetAlarmFaultObjectListResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}
