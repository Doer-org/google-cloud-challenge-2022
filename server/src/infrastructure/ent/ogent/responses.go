// Code generated by ent, DO NOT EDIT.

package ogent

import "github.com/Doer-org/google-cloud-challenge-2022/infrastructure/ent"

func NewAuthStatesCreate(e *ent.AuthStates) *AuthStatesCreate {
	if e == nil {
		return nil
	}
	var ret AuthStatesCreate
	ret.ID = e.ID
	ret.State = e.State
	ret.RedirectURL = NewOptString(e.RedirectURL)
	return &ret
}

func NewAuthStatesCreates(es []*ent.AuthStates) []AuthStatesCreate {
	if len(es) == 0 {
		return nil
	}
	r := make([]AuthStatesCreate, len(es))
	for i, e := range es {
		r[i] = NewAuthStatesCreate(e).Elem()
	}
	return r
}

func (as *AuthStatesCreate) Elem() AuthStatesCreate {
	if as == nil {
		return AuthStatesCreate{}
	}
	return *as
}

func NewAuthStatesList(e *ent.AuthStates) *AuthStatesList {
	if e == nil {
		return nil
	}
	var ret AuthStatesList
	ret.ID = e.ID
	ret.State = e.State
	ret.RedirectURL = NewOptString(e.RedirectURL)
	return &ret
}

func NewAuthStatesLists(es []*ent.AuthStates) []AuthStatesList {
	if len(es) == 0 {
		return nil
	}
	r := make([]AuthStatesList, len(es))
	for i, e := range es {
		r[i] = NewAuthStatesList(e).Elem()
	}
	return r
}

func (as *AuthStatesList) Elem() AuthStatesList {
	if as == nil {
		return AuthStatesList{}
	}
	return *as
}

func NewAuthStatesRead(e *ent.AuthStates) *AuthStatesRead {
	if e == nil {
		return nil
	}
	var ret AuthStatesRead
	ret.ID = e.ID
	ret.State = e.State
	ret.RedirectURL = NewOptString(e.RedirectURL)
	return &ret
}

func NewAuthStatesReads(es []*ent.AuthStates) []AuthStatesRead {
	if len(es) == 0 {
		return nil
	}
	r := make([]AuthStatesRead, len(es))
	for i, e := range es {
		r[i] = NewAuthStatesRead(e).Elem()
	}
	return r
}

func (as *AuthStatesRead) Elem() AuthStatesRead {
	if as == nil {
		return AuthStatesRead{}
	}
	return *as
}

func NewAuthStatesUpdate(e *ent.AuthStates) *AuthStatesUpdate {
	if e == nil {
		return nil
	}
	var ret AuthStatesUpdate
	ret.ID = e.ID
	ret.State = e.State
	ret.RedirectURL = NewOptString(e.RedirectURL)
	return &ret
}

func NewAuthStatesUpdates(es []*ent.AuthStates) []AuthStatesUpdate {
	if len(es) == 0 {
		return nil
	}
	r := make([]AuthStatesUpdate, len(es))
	for i, e := range es {
		r[i] = NewAuthStatesUpdate(e).Elem()
	}
	return r
}

func (as *AuthStatesUpdate) Elem() AuthStatesUpdate {
	if as == nil {
		return AuthStatesUpdate{}
	}
	return *as
}

func NewEStateCreate(e *ent.EState) *EStateCreate {
	if e == nil {
		return nil
	}
	var ret EStateCreate
	ret.ID = e.ID
	ret.Name = e.Name
	return &ret
}

func NewEStateCreates(es []*ent.EState) []EStateCreate {
	if len(es) == 0 {
		return nil
	}
	r := make([]EStateCreate, len(es))
	for i, e := range es {
		r[i] = NewEStateCreate(e).Elem()
	}
	return r
}

func (e *EStateCreate) Elem() EStateCreate {
	if e == nil {
		return EStateCreate{}
	}
	return *e
}

func NewEStateList(e *ent.EState) *EStateList {
	if e == nil {
		return nil
	}
	var ret EStateList
	ret.ID = e.ID
	ret.Name = e.Name
	return &ret
}

func NewEStateLists(es []*ent.EState) []EStateList {
	if len(es) == 0 {
		return nil
	}
	r := make([]EStateList, len(es))
	for i, e := range es {
		r[i] = NewEStateList(e).Elem()
	}
	return r
}

func (e *EStateList) Elem() EStateList {
	if e == nil {
		return EStateList{}
	}
	return *e
}

func NewEStateRead(e *ent.EState) *EStateRead {
	if e == nil {
		return nil
	}
	var ret EStateRead
	ret.ID = e.ID
	ret.Name = e.Name
	return &ret
}

func NewEStateReads(es []*ent.EState) []EStateRead {
	if len(es) == 0 {
		return nil
	}
	r := make([]EStateRead, len(es))
	for i, e := range es {
		r[i] = NewEStateRead(e).Elem()
	}
	return r
}

func (e *EStateRead) Elem() EStateRead {
	if e == nil {
		return EStateRead{}
	}
	return *e
}

func NewEStateUpdate(e *ent.EState) *EStateUpdate {
	if e == nil {
		return nil
	}
	var ret EStateUpdate
	ret.ID = e.ID
	ret.Name = e.Name
	return &ret
}

func NewEStateUpdates(es []*ent.EState) []EStateUpdate {
	if len(es) == 0 {
		return nil
	}
	r := make([]EStateUpdate, len(es))
	for i, e := range es {
		r[i] = NewEStateUpdate(e).Elem()
	}
	return r
}

func (e *EStateUpdate) Elem() EStateUpdate {
	if e == nil {
		return EStateUpdate{}
	}
	return *e
}

func NewEStateEventRead(e *ent.Event) *EStateEventRead {
	if e == nil {
		return nil
	}
	var ret EStateEventRead
	ret.ID = e.ID
	ret.Name = e.Name
	ret.Detail = NewOptString(e.Detail)
	ret.Location = NewOptString(e.Location)
	return &ret
}

func NewEStateEventReads(es []*ent.Event) []EStateEventRead {
	if len(es) == 0 {
		return nil
	}
	r := make([]EStateEventRead, len(es))
	for i, e := range es {
		r[i] = NewEStateEventRead(e).Elem()
	}
	return r
}

func (e *EStateEventRead) Elem() EStateEventRead {
	if e == nil {
		return EStateEventRead{}
	}
	return *e
}

func NewETypeCreate(e *ent.EType) *ETypeCreate {
	if e == nil {
		return nil
	}
	var ret ETypeCreate
	ret.ID = e.ID
	ret.Name = e.Name
	return &ret
}

func NewETypeCreates(es []*ent.EType) []ETypeCreate {
	if len(es) == 0 {
		return nil
	}
	r := make([]ETypeCreate, len(es))
	for i, e := range es {
		r[i] = NewETypeCreate(e).Elem()
	}
	return r
}

func (e *ETypeCreate) Elem() ETypeCreate {
	if e == nil {
		return ETypeCreate{}
	}
	return *e
}

func NewETypeList(e *ent.EType) *ETypeList {
	if e == nil {
		return nil
	}
	var ret ETypeList
	ret.ID = e.ID
	ret.Name = e.Name
	return &ret
}

func NewETypeLists(es []*ent.EType) []ETypeList {
	if len(es) == 0 {
		return nil
	}
	r := make([]ETypeList, len(es))
	for i, e := range es {
		r[i] = NewETypeList(e).Elem()
	}
	return r
}

func (e *ETypeList) Elem() ETypeList {
	if e == nil {
		return ETypeList{}
	}
	return *e
}

func NewETypeRead(e *ent.EType) *ETypeRead {
	if e == nil {
		return nil
	}
	var ret ETypeRead
	ret.ID = e.ID
	ret.Name = e.Name
	return &ret
}

func NewETypeReads(es []*ent.EType) []ETypeRead {
	if len(es) == 0 {
		return nil
	}
	r := make([]ETypeRead, len(es))
	for i, e := range es {
		r[i] = NewETypeRead(e).Elem()
	}
	return r
}

func (e *ETypeRead) Elem() ETypeRead {
	if e == nil {
		return ETypeRead{}
	}
	return *e
}

func NewETypeUpdate(e *ent.EType) *ETypeUpdate {
	if e == nil {
		return nil
	}
	var ret ETypeUpdate
	ret.ID = e.ID
	ret.Name = e.Name
	return &ret
}

func NewETypeUpdates(es []*ent.EType) []ETypeUpdate {
	if len(es) == 0 {
		return nil
	}
	r := make([]ETypeUpdate, len(es))
	for i, e := range es {
		r[i] = NewETypeUpdate(e).Elem()
	}
	return r
}

func (e *ETypeUpdate) Elem() ETypeUpdate {
	if e == nil {
		return ETypeUpdate{}
	}
	return *e
}

func NewETypeEventRead(e *ent.Event) *ETypeEventRead {
	if e == nil {
		return nil
	}
	var ret ETypeEventRead
	ret.ID = e.ID
	ret.Name = e.Name
	ret.Detail = NewOptString(e.Detail)
	ret.Location = NewOptString(e.Location)
	return &ret
}

func NewETypeEventReads(es []*ent.Event) []ETypeEventRead {
	if len(es) == 0 {
		return nil
	}
	r := make([]ETypeEventRead, len(es))
	for i, e := range es {
		r[i] = NewETypeEventRead(e).Elem()
	}
	return r
}

func (e *ETypeEventRead) Elem() ETypeEventRead {
	if e == nil {
		return ETypeEventRead{}
	}
	return *e
}

func NewEcommentCreate(e *ent.Ecomment) *EcommentCreate {
	if e == nil {
		return nil
	}
	var ret EcommentCreate
	ret.ID = e.ID
	ret.Body = e.Body
	return &ret
}

func NewEcommentCreates(es []*ent.Ecomment) []EcommentCreate {
	if len(es) == 0 {
		return nil
	}
	r := make([]EcommentCreate, len(es))
	for i, e := range es {
		r[i] = NewEcommentCreate(e).Elem()
	}
	return r
}

func (e *EcommentCreate) Elem() EcommentCreate {
	if e == nil {
		return EcommentCreate{}
	}
	return *e
}

func NewEcommentList(e *ent.Ecomment) *EcommentList {
	if e == nil {
		return nil
	}
	var ret EcommentList
	ret.ID = e.ID
	ret.Body = e.Body
	return &ret
}

func NewEcommentLists(es []*ent.Ecomment) []EcommentList {
	if len(es) == 0 {
		return nil
	}
	r := make([]EcommentList, len(es))
	for i, e := range es {
		r[i] = NewEcommentList(e).Elem()
	}
	return r
}

func (e *EcommentList) Elem() EcommentList {
	if e == nil {
		return EcommentList{}
	}
	return *e
}

func NewEcommentRead(e *ent.Ecomment) *EcommentRead {
	if e == nil {
		return nil
	}
	var ret EcommentRead
	ret.ID = e.ID
	ret.Body = e.Body
	return &ret
}

func NewEcommentReads(es []*ent.Ecomment) []EcommentRead {
	if len(es) == 0 {
		return nil
	}
	r := make([]EcommentRead, len(es))
	for i, e := range es {
		r[i] = NewEcommentRead(e).Elem()
	}
	return r
}

func (e *EcommentRead) Elem() EcommentRead {
	if e == nil {
		return EcommentRead{}
	}
	return *e
}

func NewEcommentUpdate(e *ent.Ecomment) *EcommentUpdate {
	if e == nil {
		return nil
	}
	var ret EcommentUpdate
	ret.ID = e.ID
	ret.Body = e.Body
	return &ret
}

func NewEcommentUpdates(es []*ent.Ecomment) []EcommentUpdate {
	if len(es) == 0 {
		return nil
	}
	r := make([]EcommentUpdate, len(es))
	for i, e := range es {
		r[i] = NewEcommentUpdate(e).Elem()
	}
	return r
}

func (e *EcommentUpdate) Elem() EcommentUpdate {
	if e == nil {
		return EcommentUpdate{}
	}
	return *e
}

func NewEcommentEventRead(e *ent.Event) *EcommentEventRead {
	if e == nil {
		return nil
	}
	var ret EcommentEventRead
	ret.ID = e.ID
	ret.Name = e.Name
	ret.Detail = NewOptString(e.Detail)
	ret.Location = NewOptString(e.Location)
	return &ret
}

func NewEcommentEventReads(es []*ent.Event) []EcommentEventRead {
	if len(es) == 0 {
		return nil
	}
	r := make([]EcommentEventRead, len(es))
	for i, e := range es {
		r[i] = NewEcommentEventRead(e).Elem()
	}
	return r
}

func (e *EcommentEventRead) Elem() EcommentEventRead {
	if e == nil {
		return EcommentEventRead{}
	}
	return *e
}

func NewEcommentUserRead(e *ent.User) *EcommentUserRead {
	if e == nil {
		return nil
	}
	var ret EcommentUserRead
	ret.ID = e.ID
	ret.Age = NewOptInt(e.Age)
	ret.Name = e.Name
	ret.Authenticated = e.Authenticated
	ret.Mail = NewOptString(e.Mail)
	ret.Icon = e.Icon
	return &ret
}

func NewEcommentUserReads(es []*ent.User) []EcommentUserRead {
	if len(es) == 0 {
		return nil
	}
	r := make([]EcommentUserRead, len(es))
	for i, e := range es {
		r[i] = NewEcommentUserRead(e).Elem()
	}
	return r
}

func (u *EcommentUserRead) Elem() EcommentUserRead {
	if u == nil {
		return EcommentUserRead{}
	}
	return *u
}

func NewEventCreate(e *ent.Event) *EventCreate {
	if e == nil {
		return nil
	}
	var ret EventCreate
	ret.ID = e.ID
	ret.Name = e.Name
	ret.Detail = NewOptString(e.Detail)
	ret.Location = NewOptString(e.Location)
	return &ret
}

func NewEventCreates(es []*ent.Event) []EventCreate {
	if len(es) == 0 {
		return nil
	}
	r := make([]EventCreate, len(es))
	for i, e := range es {
		r[i] = NewEventCreate(e).Elem()
	}
	return r
}

func (e *EventCreate) Elem() EventCreate {
	if e == nil {
		return EventCreate{}
	}
	return *e
}

func NewEventList(e *ent.Event) *EventList {
	if e == nil {
		return nil
	}
	var ret EventList
	ret.ID = e.ID
	ret.Name = e.Name
	ret.Detail = NewOptString(e.Detail)
	ret.Location = NewOptString(e.Location)
	return &ret
}

func NewEventLists(es []*ent.Event) []EventList {
	if len(es) == 0 {
		return nil
	}
	r := make([]EventList, len(es))
	for i, e := range es {
		r[i] = NewEventList(e).Elem()
	}
	return r
}

func (e *EventList) Elem() EventList {
	if e == nil {
		return EventList{}
	}
	return *e
}

func NewEventRead(e *ent.Event) *EventRead {
	if e == nil {
		return nil
	}
	var ret EventRead
	ret.ID = e.ID
	ret.Name = e.Name
	ret.Detail = NewOptString(e.Detail)
	ret.Location = NewOptString(e.Location)
	return &ret
}

func NewEventReads(es []*ent.Event) []EventRead {
	if len(es) == 0 {
		return nil
	}
	r := make([]EventRead, len(es))
	for i, e := range es {
		r[i] = NewEventRead(e).Elem()
	}
	return r
}

func (e *EventRead) Elem() EventRead {
	if e == nil {
		return EventRead{}
	}
	return *e
}

func NewEventUpdate(e *ent.Event) *EventUpdate {
	if e == nil {
		return nil
	}
	var ret EventUpdate
	ret.ID = e.ID
	ret.Name = e.Name
	ret.Detail = NewOptString(e.Detail)
	ret.Location = NewOptString(e.Location)
	return &ret
}

func NewEventUpdates(es []*ent.Event) []EventUpdate {
	if len(es) == 0 {
		return nil
	}
	r := make([]EventUpdate, len(es))
	for i, e := range es {
		r[i] = NewEventUpdate(e).Elem()
	}
	return r
}

func (e *EventUpdate) Elem() EventUpdate {
	if e == nil {
		return EventUpdate{}
	}
	return *e
}

func NewEventStateRead(e *ent.EState) *EventStateRead {
	if e == nil {
		return nil
	}
	var ret EventStateRead
	ret.ID = e.ID
	ret.Name = e.Name
	return &ret
}

func NewEventStateReads(es []*ent.EState) []EventStateRead {
	if len(es) == 0 {
		return nil
	}
	r := make([]EventStateRead, len(es))
	for i, e := range es {
		r[i] = NewEventStateRead(e).Elem()
	}
	return r
}

func (e *EventStateRead) Elem() EventStateRead {
	if e == nil {
		return EventStateRead{}
	}
	return *e
}

func NewEventTypeRead(e *ent.EType) *EventTypeRead {
	if e == nil {
		return nil
	}
	var ret EventTypeRead
	ret.ID = e.ID
	ret.Name = e.Name
	return &ret
}

func NewEventTypeReads(es []*ent.EType) []EventTypeRead {
	if len(es) == 0 {
		return nil
	}
	r := make([]EventTypeRead, len(es))
	for i, e := range es {
		r[i] = NewEventTypeRead(e).Elem()
	}
	return r
}

func (e *EventTypeRead) Elem() EventTypeRead {
	if e == nil {
		return EventTypeRead{}
	}
	return *e
}

func NewEventUsersList(e *ent.User) *EventUsersList {
	if e == nil {
		return nil
	}
	var ret EventUsersList
	ret.ID = e.ID
	ret.Age = NewOptInt(e.Age)
	ret.Name = e.Name
	ret.Authenticated = e.Authenticated
	ret.Mail = NewOptString(e.Mail)
	ret.Icon = e.Icon
	return &ret
}

func NewEventUsersLists(es []*ent.User) []EventUsersList {
	if len(es) == 0 {
		return nil
	}
	r := make([]EventUsersList, len(es))
	for i, e := range es {
		r[i] = NewEventUsersList(e).Elem()
	}
	return r
}

func (u *EventUsersList) Elem() EventUsersList {
	if u == nil {
		return EventUsersList{}
	}
	return *u
}

func NewGoogleAuthCreate(e *ent.GoogleAuth) *GoogleAuthCreate {
	if e == nil {
		return nil
	}
	var ret GoogleAuthCreate
	ret.ID = e.ID
	ret.UserID = e.UserID
	ret.AccessToken = e.AccessToken
	ret.RefreshToken = e.RefreshToken
	ret.Expiry = e.Expiry
	return &ret
}

func NewGoogleAuthCreates(es []*ent.GoogleAuth) []GoogleAuthCreate {
	if len(es) == 0 {
		return nil
	}
	r := make([]GoogleAuthCreate, len(es))
	for i, e := range es {
		r[i] = NewGoogleAuthCreate(e).Elem()
	}
	return r
}

func (ga *GoogleAuthCreate) Elem() GoogleAuthCreate {
	if ga == nil {
		return GoogleAuthCreate{}
	}
	return *ga
}

func NewGoogleAuthList(e *ent.GoogleAuth) *GoogleAuthList {
	if e == nil {
		return nil
	}
	var ret GoogleAuthList
	ret.ID = e.ID
	ret.UserID = e.UserID
	ret.AccessToken = e.AccessToken
	ret.RefreshToken = e.RefreshToken
	ret.Expiry = e.Expiry
	return &ret
}

func NewGoogleAuthLists(es []*ent.GoogleAuth) []GoogleAuthList {
	if len(es) == 0 {
		return nil
	}
	r := make([]GoogleAuthList, len(es))
	for i, e := range es {
		r[i] = NewGoogleAuthList(e).Elem()
	}
	return r
}

func (ga *GoogleAuthList) Elem() GoogleAuthList {
	if ga == nil {
		return GoogleAuthList{}
	}
	return *ga
}

func NewGoogleAuthRead(e *ent.GoogleAuth) *GoogleAuthRead {
	if e == nil {
		return nil
	}
	var ret GoogleAuthRead
	ret.ID = e.ID
	ret.UserID = e.UserID
	ret.AccessToken = e.AccessToken
	ret.RefreshToken = e.RefreshToken
	ret.Expiry = e.Expiry
	return &ret
}

func NewGoogleAuthReads(es []*ent.GoogleAuth) []GoogleAuthRead {
	if len(es) == 0 {
		return nil
	}
	r := make([]GoogleAuthRead, len(es))
	for i, e := range es {
		r[i] = NewGoogleAuthRead(e).Elem()
	}
	return r
}

func (ga *GoogleAuthRead) Elem() GoogleAuthRead {
	if ga == nil {
		return GoogleAuthRead{}
	}
	return *ga
}

func NewGoogleAuthUpdate(e *ent.GoogleAuth) *GoogleAuthUpdate {
	if e == nil {
		return nil
	}
	var ret GoogleAuthUpdate
	ret.ID = e.ID
	ret.UserID = e.UserID
	ret.AccessToken = e.AccessToken
	ret.RefreshToken = e.RefreshToken
	ret.Expiry = e.Expiry
	return &ret
}

func NewGoogleAuthUpdates(es []*ent.GoogleAuth) []GoogleAuthUpdate {
	if len(es) == 0 {
		return nil
	}
	r := make([]GoogleAuthUpdate, len(es))
	for i, e := range es {
		r[i] = NewGoogleAuthUpdate(e).Elem()
	}
	return r
}

func (ga *GoogleAuthUpdate) Elem() GoogleAuthUpdate {
	if ga == nil {
		return GoogleAuthUpdate{}
	}
	return *ga
}

func NewGoogleAuthUserRead(e *ent.User) *GoogleAuthUserRead {
	if e == nil {
		return nil
	}
	var ret GoogleAuthUserRead
	ret.ID = e.ID
	ret.Age = NewOptInt(e.Age)
	ret.Name = e.Name
	ret.Authenticated = e.Authenticated
	ret.Mail = NewOptString(e.Mail)
	ret.Icon = e.Icon
	return &ret
}

func NewGoogleAuthUserReads(es []*ent.User) []GoogleAuthUserRead {
	if len(es) == 0 {
		return nil
	}
	r := make([]GoogleAuthUserRead, len(es))
	for i, e := range es {
		r[i] = NewGoogleAuthUserRead(e).Elem()
	}
	return r
}

func (u *GoogleAuthUserRead) Elem() GoogleAuthUserRead {
	if u == nil {
		return GoogleAuthUserRead{}
	}
	return *u
}

func NewLoginSessionsCreate(e *ent.LoginSessions) *LoginSessionsCreate {
	if e == nil {
		return nil
	}
	var ret LoginSessionsCreate
	ret.ID = e.ID
	ret.UserID = e.UserID
	return &ret
}

func NewLoginSessionsCreates(es []*ent.LoginSessions) []LoginSessionsCreate {
	if len(es) == 0 {
		return nil
	}
	r := make([]LoginSessionsCreate, len(es))
	for i, e := range es {
		r[i] = NewLoginSessionsCreate(e).Elem()
	}
	return r
}

func (ls *LoginSessionsCreate) Elem() LoginSessionsCreate {
	if ls == nil {
		return LoginSessionsCreate{}
	}
	return *ls
}

func NewLoginSessionsList(e *ent.LoginSessions) *LoginSessionsList {
	if e == nil {
		return nil
	}
	var ret LoginSessionsList
	ret.ID = e.ID
	ret.UserID = e.UserID
	return &ret
}

func NewLoginSessionsLists(es []*ent.LoginSessions) []LoginSessionsList {
	if len(es) == 0 {
		return nil
	}
	r := make([]LoginSessionsList, len(es))
	for i, e := range es {
		r[i] = NewLoginSessionsList(e).Elem()
	}
	return r
}

func (ls *LoginSessionsList) Elem() LoginSessionsList {
	if ls == nil {
		return LoginSessionsList{}
	}
	return *ls
}

func NewLoginSessionsRead(e *ent.LoginSessions) *LoginSessionsRead {
	if e == nil {
		return nil
	}
	var ret LoginSessionsRead
	ret.ID = e.ID
	ret.UserID = e.UserID
	return &ret
}

func NewLoginSessionsReads(es []*ent.LoginSessions) []LoginSessionsRead {
	if len(es) == 0 {
		return nil
	}
	r := make([]LoginSessionsRead, len(es))
	for i, e := range es {
		r[i] = NewLoginSessionsRead(e).Elem()
	}
	return r
}

func (ls *LoginSessionsRead) Elem() LoginSessionsRead {
	if ls == nil {
		return LoginSessionsRead{}
	}
	return *ls
}

func NewLoginSessionsUpdate(e *ent.LoginSessions) *LoginSessionsUpdate {
	if e == nil {
		return nil
	}
	var ret LoginSessionsUpdate
	ret.ID = e.ID
	ret.UserID = e.UserID
	return &ret
}

func NewLoginSessionsUpdates(es []*ent.LoginSessions) []LoginSessionsUpdate {
	if len(es) == 0 {
		return nil
	}
	r := make([]LoginSessionsUpdate, len(es))
	for i, e := range es {
		r[i] = NewLoginSessionsUpdate(e).Elem()
	}
	return r
}

func (ls *LoginSessionsUpdate) Elem() LoginSessionsUpdate {
	if ls == nil {
		return LoginSessionsUpdate{}
	}
	return *ls
}

func NewLoginSessionsUserRead(e *ent.User) *LoginSessionsUserRead {
	if e == nil {
		return nil
	}
	var ret LoginSessionsUserRead
	ret.ID = e.ID
	ret.Age = NewOptInt(e.Age)
	ret.Name = e.Name
	ret.Authenticated = e.Authenticated
	ret.Mail = NewOptString(e.Mail)
	ret.Icon = e.Icon
	return &ret
}

func NewLoginSessionsUserReads(es []*ent.User) []LoginSessionsUserRead {
	if len(es) == 0 {
		return nil
	}
	r := make([]LoginSessionsUserRead, len(es))
	for i, e := range es {
		r[i] = NewLoginSessionsUserRead(e).Elem()
	}
	return r
}

func (u *LoginSessionsUserRead) Elem() LoginSessionsUserRead {
	if u == nil {
		return LoginSessionsUserRead{}
	}
	return *u
}

func NewUserCreate(e *ent.User) *UserCreate {
	if e == nil {
		return nil
	}
	var ret UserCreate
	ret.ID = e.ID
	ret.Age = NewOptInt(e.Age)
	ret.Name = e.Name
	ret.Authenticated = e.Authenticated
	ret.Mail = NewOptString(e.Mail)
	ret.Icon = e.Icon
	return &ret
}

func NewUserCreates(es []*ent.User) []UserCreate {
	if len(es) == 0 {
		return nil
	}
	r := make([]UserCreate, len(es))
	for i, e := range es {
		r[i] = NewUserCreate(e).Elem()
	}
	return r
}

func (u *UserCreate) Elem() UserCreate {
	if u == nil {
		return UserCreate{}
	}
	return *u
}

func NewUserList(e *ent.User) *UserList {
	if e == nil {
		return nil
	}
	var ret UserList
	ret.ID = e.ID
	ret.Age = NewOptInt(e.Age)
	ret.Name = e.Name
	ret.Authenticated = e.Authenticated
	ret.Mail = NewOptString(e.Mail)
	ret.Icon = e.Icon
	return &ret
}

func NewUserLists(es []*ent.User) []UserList {
	if len(es) == 0 {
		return nil
	}
	r := make([]UserList, len(es))
	for i, e := range es {
		r[i] = NewUserList(e).Elem()
	}
	return r
}

func (u *UserList) Elem() UserList {
	if u == nil {
		return UserList{}
	}
	return *u
}

func NewUserRead(e *ent.User) *UserRead {
	if e == nil {
		return nil
	}
	var ret UserRead
	ret.ID = e.ID
	ret.Age = NewOptInt(e.Age)
	ret.Name = e.Name
	ret.Authenticated = e.Authenticated
	ret.Mail = NewOptString(e.Mail)
	ret.Icon = e.Icon
	return &ret
}

func NewUserReads(es []*ent.User) []UserRead {
	if len(es) == 0 {
		return nil
	}
	r := make([]UserRead, len(es))
	for i, e := range es {
		r[i] = NewUserRead(e).Elem()
	}
	return r
}

func (u *UserRead) Elem() UserRead {
	if u == nil {
		return UserRead{}
	}
	return *u
}

func NewUserUpdate(e *ent.User) *UserUpdate {
	if e == nil {
		return nil
	}
	var ret UserUpdate
	ret.ID = e.ID
	ret.Age = NewOptInt(e.Age)
	ret.Name = e.Name
	ret.Authenticated = e.Authenticated
	ret.Mail = NewOptString(e.Mail)
	ret.Icon = e.Icon
	return &ret
}

func NewUserUpdates(es []*ent.User) []UserUpdate {
	if len(es) == 0 {
		return nil
	}
	r := make([]UserUpdate, len(es))
	for i, e := range es {
		r[i] = NewUserUpdate(e).Elem()
	}
	return r
}

func (u *UserUpdate) Elem() UserUpdate {
	if u == nil {
		return UserUpdate{}
	}
	return *u
}

func NewUserEventsList(e *ent.Event) *UserEventsList {
	if e == nil {
		return nil
	}
	var ret UserEventsList
	ret.ID = e.ID
	ret.Name = e.Name
	ret.Detail = NewOptString(e.Detail)
	ret.Location = NewOptString(e.Location)
	return &ret
}

func NewUserEventsLists(es []*ent.Event) []UserEventsList {
	if len(es) == 0 {
		return nil
	}
	r := make([]UserEventsList, len(es))
	for i, e := range es {
		r[i] = NewUserEventsList(e).Elem()
	}
	return r
}

func (e *UserEventsList) Elem() UserEventsList {
	if e == nil {
		return UserEventsList{}
	}
	return *e
}
