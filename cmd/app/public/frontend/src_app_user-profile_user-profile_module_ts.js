"use strict";
(self["webpackChunkshopping_lists_and_recipes"] = self["webpackChunkshopping_lists_and_recipes"] || []).push([["src_app_user-profile_user-profile_module_ts"],{

/***/ 9488:
/*!**************************************************!*\
  !*** ./src/app/shared/secure-image-pipe.pipe.ts ***!
  \**************************************************/
/***/ ((__unused_webpack_module, __webpack_exports__, __webpack_require__) => {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   "SecureImagePipe": () => (/* binding */ SecureImagePipe)
/* harmony export */ });
/* harmony import */ var tslib__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! tslib */ 4929);
/* harmony import */ var _angular_common_http__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! @angular/common/http */ 8784);
/* harmony import */ var src_environments_environment__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! src/environments/environment */ 2340);
/* harmony import */ var _angular_core__WEBPACK_IMPORTED_MODULE_4__ = __webpack_require__(/*! @angular/core */ 3184);
/* harmony import */ var _auth_auth_service__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! ../auth/auth.service */ 384);






class SecureImagePipe {
    constructor(http, auth) {
        this.http = http;
        this.auth = auth;
    }
    transform(src) {
        return (0,tslib__WEBPACK_IMPORTED_MODULE_2__.__awaiter)(this, void 0, void 0, function* () {
            const token = this.auth.GetUserToken();
            const headers = new _angular_common_http__WEBPACK_IMPORTED_MODULE_3__.HttpHeaders({
                Auth: token,
                ApiKey: src_environments_environment__WEBPACK_IMPORTED_MODULE_0__.environment.ApiKey
            });
            try {
                const imageBlob = yield this.http.get(src, { headers, responseType: 'blob' }).toPromise();
                const reader = new FileReader();
                return new Promise((resolve, reject) => {
                    reader.onloadend = () => resolve(reader.result);
                    reader.readAsDataURL(imageBlob);
                });
            }
            catch (_a) {
                return '/favicon.ico';
            }
        });
    }
}
SecureImagePipe.ɵfac = function SecureImagePipe_Factory(t) { return new (t || SecureImagePipe)(_angular_core__WEBPACK_IMPORTED_MODULE_4__["ɵɵdirectiveInject"](_angular_common_http__WEBPACK_IMPORTED_MODULE_3__.HttpClient, 16), _angular_core__WEBPACK_IMPORTED_MODULE_4__["ɵɵdirectiveInject"](_auth_auth_service__WEBPACK_IMPORTED_MODULE_1__.AuthService, 16)); };
SecureImagePipe.ɵpipe = /*@__PURE__*/ _angular_core__WEBPACK_IMPORTED_MODULE_4__["ɵɵdefinePipe"]({ name: "secureImagePipe", type: SecureImagePipe, pure: true });


/***/ }),

/***/ 7960:
/*!********************************************************!*\
  !*** ./src/app/user-profile/user-profile.component.ts ***!
  \********************************************************/
/***/ ((__unused_webpack_module, __webpack_exports__, __webpack_require__) => {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   "UserProfileComponent": () => (/* binding */ UserProfileComponent)
/* harmony export */ });
/* harmony import */ var src_environments_environment__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! src/environments/environment */ 2340);
/* harmony import */ var _angular_core__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! @angular/core */ 3184);
/* harmony import */ var _angular_router__WEBPACK_IMPORTED_MODULE_4__ = __webpack_require__(/*! @angular/router */ 2816);
/* harmony import */ var _shared_data_storage_service__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! ../shared/data-storage.service */ 3649);
/* harmony import */ var _angular_common__WEBPACK_IMPORTED_MODULE_5__ = __webpack_require__(/*! @angular/common */ 6362);
/* harmony import */ var _ng_bootstrap_ng_bootstrap__WEBPACK_IMPORTED_MODULE_6__ = __webpack_require__(/*! @ng-bootstrap/ng-bootstrap */ 7544);
/* harmony import */ var _angular_forms__WEBPACK_IMPORTED_MODULE_7__ = __webpack_require__(/*! @angular/forms */ 587);
/* harmony import */ var _shared_secure_image_pipe_pipe__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! ../shared/secure-image-pipe.pipe */ 9488);








function UserProfileComponent_div_1_Template(rf, ctx) { if (rf & 1) {
    _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵelementStart"](0, "div", 5)(1, "span", 6);
    _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵtext"](2, "Loading...");
    _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵelementEnd"]()();
} }
function UserProfileComponent_ngb_alert_2_Template(rf, ctx) { if (rf & 1) {
    const _r6 = _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵgetCurrentView"]();
    _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵelementStart"](0, "ngb-alert", 7);
    _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵlistener"]("close", function UserProfileComponent_ngb_alert_2_Template_ngb_alert_close_0_listener() { _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵrestoreView"](_r6); const ctx_r5 = _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵnextContext"](); return ctx_r5.ShowMessage = false; });
    _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵtext"](1);
    _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵelementEnd"]();
} if (rf & 2) {
    const ctx_r1 = _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵnextContext"]();
    _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵproperty"]("type", ctx_r1.MessageType);
    _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵadvance"](1);
    _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵtextInterpolate1"](" ", ctx_r1.ResponseFromBackend.Error.Message, "");
} }
function UserProfileComponent_form_3_Template(rf, ctx) { if (rf & 1) {
    const _r9 = _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵgetCurrentView"]();
    _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵelementStart"](0, "form", 8, 9);
    _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵlistener"]("ngSubmit", function UserProfileComponent_form_3_Template_form_ngSubmit_0_listener() { _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵrestoreView"](_r9); const _r7 = _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵreference"](1); const ctx_r8 = _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵnextContext"](); return ctx_r8.OnSaveClick(_r7); });
    _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵelementStart"](2, "div", 10)(3, "div", 11)(4, "h2");
    _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵtext"](5, "Profile");
    _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵelementEnd"]();
    _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵelementStart"](6, "div", 10)(7, "label", 12);
    _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵtext"](8, "Name");
    _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵelementEnd"]();
    _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵelement"](9, "input", 13);
    _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵelementEnd"]();
    _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵelementStart"](10, "div", 10)(11, "label", 14);
    _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵtext"](12, "Email");
    _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵelementEnd"]();
    _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵelement"](13, "input", 15);
    _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵelementEnd"]();
    _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵelementStart"](14, "div", 10)(15, "label", 16);
    _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵtext"](16, "Phone");
    _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵelementEnd"]();
    _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵelement"](17, "input", 17);
    _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵelementEnd"]();
    _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵelementStart"](18, "label", 18);
    _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵtext"](19, "Password");
    _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵelementEnd"]();
    _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵelementStart"](20, "div", 19)(21, "div", 20)(22, "div", 21)(23, "input", 22);
    _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵlistener"]("ngModelChange", function UserProfileComponent_form_3_Template_input_ngModelChange_23_listener($event) { _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵrestoreView"](_r9); const ctx_r10 = _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵnextContext"](); return ctx_r10.changepassword = $event; });
    _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵelementEnd"]()()();
    _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵelement"](24, "input", 23);
    _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵelementEnd"]();
    _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵelementStart"](25, "button", 24);
    _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵtext"](26, "Save");
    _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵelementEnd"]();
    _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵelementStart"](27, "button", 25);
    _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵtext"](28, "Cancel");
    _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵelementEnd"]()()()();
} if (rf & 2) {
    const _r7 = _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵreference"](1);
    const ctx_r2 = _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵnextContext"]();
    _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵadvance"](9);
    _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵproperty"]("ngModel", ctx_r2.UserToEdit.Name);
    _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵadvance"](4);
    _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵproperty"]("ngModel", ctx_r2.UserToEdit.Email);
    _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵadvance"](4);
    _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵproperty"]("ngModel", ctx_r2.UserToEdit.Phone);
    _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵadvance"](6);
    _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵproperty"]("ngModel", ctx_r2.changepassword);
    _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵadvance"](1);
    _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵproperty"]("disabled", !ctx_r2.changepassword);
    _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵadvance"](1);
    _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵproperty"]("disabled", _r7.invalid);
} }
function UserProfileComponent_div_4_Template(rf, ctx) { if (rf & 1) {
    const _r12 = _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵgetCurrentView"]();
    _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵelementStart"](0, "div", 10)(1, "h2");
    _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵtext"](2, "Second factor enabled");
    _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵelementEnd"]();
    _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵelementStart"](3, "button", 26);
    _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵlistener"]("click", function UserProfileComponent_div_4_Template_button_click_3_listener() { _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵrestoreView"](_r12); const ctx_r11 = _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵnextContext"](); return ctx_r11.OnUnlinkTwoFactor(); });
    _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵtext"](4, "Disable");
    _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵelementEnd"]()();
} }
function UserProfileComponent_form_5_Template(rf, ctx) { if (rf & 1) {
    const _r15 = _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵgetCurrentView"]();
    _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵelementStart"](0, "form", 8, 27);
    _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵlistener"]("ngSubmit", function UserProfileComponent_form_5_Template_form_ngSubmit_0_listener() { _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵrestoreView"](_r15); const _r13 = _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵreference"](1); const ctx_r14 = _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵnextContext"](); return ctx_r14.OnLinkTwoFactor(_r13); });
    _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵelementStart"](2, "div", 10)(3, "div", 11)(4, "h2");
    _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵtext"](5, "Set second factor");
    _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵelementEnd"]();
    _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵelementStart"](6, "h4");
    _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵtext"](7, "1. Scan with ");
    _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵelementStart"](8, "a", 28);
    _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵtext"](9, "Authenticator");
    _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵelementEnd"]();
    _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵtext"](10, ": ");
    _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵelementEnd"]();
    _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵelement"](11, "img", 29);
    _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵpipe"](12, "async");
    _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵpipe"](13, "secureImagePipe");
    _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵelementStart"](14, "h4");
    _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵtext"](15, "2. Enter the token:");
    _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵelementEnd"]();
    _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵelementStart"](16, "label", 18);
    _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵtext"](17, "Token");
    _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵelementEnd"]();
    _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵelementStart"](18, "div", 19)(19, "div", 20)(20, "div", 21)(21, "input", 30);
    _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵlistener"]("ngModelChange", function UserProfileComponent_form_5_Template_input_ngModelChange_21_listener($event) { _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵrestoreView"](_r15); const ctx_r16 = _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵnextContext"](); return ctx_r16.SetTwoFactor = $event; });
    _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵelementEnd"]()()();
    _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵelement"](22, "input", 31);
    _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵelementEnd"]();
    _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵelementStart"](23, "button", 32);
    _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵtext"](24, "Save");
    _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵelementEnd"]()()()();
} if (rf & 2) {
    const ctx_r4 = _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵnextContext"]();
    _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵadvance"](8);
    _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵproperty"]("href", ctx_r4.AuthUrl, _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵsanitizeUrl"]);
    _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵadvance"](3);
    _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵproperty"]("src", _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵpipeBind1"](12, 4, _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵpipeBind1"](13, 6, ctx_r4.QrUrl)), _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵsanitizeUrl"]);
    _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵadvance"](10);
    _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵproperty"]("ngModel", ctx_r4.SetTwoFactor);
    _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵadvance"](1);
    _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵproperty"]("disabled", !ctx_r4.SetTwoFactor);
} }
class UserProfileComponent {
    constructor(activatedroute, router, datastore) {
        this.activatedroute = activatedroute;
        this.router = router;
        this.datastore = datastore;
    }
    ngOnDestroy() {
        this.FetchUser.unsubscribe();
        this.RecivedErrorSub.unsubscribe();
        this.DataLoading.unsubscribe();
        this.SaveUser.unsubscribe();
        this.LinkUnlinkTFA.unsubscribe();
    }
    ngOnInit() {
        this.AuthUrl = src_environments_environment__WEBPACK_IMPORTED_MODULE_0__.environment.GetAuthenticatorUrl;
        this.QrUrl = src_environments_environment__WEBPACK_IMPORTED_MODULE_0__.environment.GetTOTPQRCodeUrl;
        this.LinkUnlinkTFA = this.datastore.TwoFactorSub.subscribe((ThisUser) => {
            this.SetUserAndTFA(ThisUser);
        });
        this.SaveUser = this.datastore.UserUpdateInsert.subscribe((ThisUser) => {
            this.SetUserAndTFA(ThisUser);
        });
        this.FetchUser = this.datastore.CurrentUserFetch.subscribe((ThisUser) => {
            this.SetUserAndTFA(ThisUser);
        });
        this.RecivedErrorSub = this.datastore.RecivedError.subscribe((response) => {
            this.ShowMessage = true;
            this.ResponseFromBackend = response;
            setTimeout(() => this.ShowMessage = false, 5000);
            if (response) {
                switch (response.Error.Code) {
                    case 200:
                        this.MessageType = 'success';
                        break;
                    default:
                        this.MessageType = 'danger';
                        break;
                }
            }
        });
        this.DataLoading = this.datastore.LoadingData.subscribe((State) => {
            this.IsLoading = State;
        });
        this.datastore.FetchCurrentUser();
    }
    SetUserAndTFA(ThisUser) {
        this.UserToEdit = ThisUser;
        if (this.UserToEdit) {
            this.TwoFactorEnabled = this.UserToEdit.SecondFactor;
        }
    }
    OnSaveClick(SubmittedForm) {
        if (SubmittedForm.valid) {
            if (SubmittedForm.value.changepassword && SubmittedForm.value.newpassword.length === 0) {
                return;
            }
            this.UserToEdit.Email = SubmittedForm.value.useremail;
            this.UserToEdit.Name = SubmittedForm.value.username;
            this.UserToEdit.Phone = SubmittedForm.value.userphone;
            this.datastore.SaveCurrentUser(this.UserToEdit, SubmittedForm.value.changepassword, SubmittedForm.value.newpassword);
        }
    }
    OnLinkTwoFactor(SubmittedForm) {
        if (SubmittedForm.valid) {
            this.datastore.LinkTwoFactor(SubmittedForm.value.passkey, this.UserToEdit);
        }
    }
    OnUnlinkTwoFactor() {
        if (this.UserToEdit.SecondFactor) {
            this.datastore.UnlinkTwoFactor(this.UserToEdit);
        }
    }
}
UserProfileComponent.ɵfac = function UserProfileComponent_Factory(t) { return new (t || UserProfileComponent)(_angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵdirectiveInject"](_angular_router__WEBPACK_IMPORTED_MODULE_4__.ActivatedRoute), _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵdirectiveInject"](_angular_router__WEBPACK_IMPORTED_MODULE_4__.Router), _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵdirectiveInject"](_shared_data_storage_service__WEBPACK_IMPORTED_MODULE_1__.DataStorageService)); };
UserProfileComponent.ɵcmp = /*@__PURE__*/ _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵdefineComponent"]({ type: UserProfileComponent, selectors: [["app-user-profile"]], decls: 6, vars: 5, consts: [[1, "text-center"], ["class", "spinner-border", "role", "status", 4, "ngIf"], [3, "type", "close", 4, "ngIf"], [3, "ngSubmit", 4, "ngIf"], ["class", "form-group", 4, "ngIf"], ["role", "status", 1, "spinner-border"], [1, "sr-only"], [3, "type", "close"], [3, "ngSubmit"], ["UserProfileForm", "ngForm"], [1, "form-group"], [2, "margin", "3px"], ["for", "name"], ["type", "text", "id", "name", "placeholder", "example", "name", "username", "required", "", 1, "form-control", "mb-1", 3, "ngModel"], ["for", "email"], ["type", "email", "id", "email", "placeholder", "exampe@example.com", "name", "useremail", "required", "", "email", "", 1, "form-control", "mb-1", 3, "ngModel"], ["for", "phone"], ["type", "tel", "id", "phone", "placeholder", "+7 (965) 777-77-77", "name", "userphone", "pattern", "^((8|\\+7)[\\- ]?)?(\\(?\\d{3,4}\\)?[\\- ]?)?[\\d\\- ]{5,10}$", 1, "form-control", "mb-1", 3, "ngModel"], ["for", "newpassword"], [1, "input-group", "mb-3"], [1, "input-group-prepend"], [1, "input-group-text"], ["type", "checkbox", "id", "changepassword", "name", "changepassword", 3, "ngModel", "ngModelChange"], ["type", "password", "id", "newpassword", "name", "newpassword", "placeholder", "Enter new password", "ngModel", "", 1, "form-control", 3, "disabled"], ["type", "submit", 1, "btn", "btn-outline-primary", 3, "disabled"], ["type", "button", "routerLink", "..", 1, "btn", "btn-outline-danger", 2, "margin-left", "3px"], ["type", "button", 1, "btn", "btn-outline-danger", 3, "click"], ["UserTwoFactorForm", "ngForm"], [3, "href"], ["alt", "QR code", 3, "src"], ["type", "checkbox", "id", "enabletwofactor", "name", "enabletwofactor", "required", "", 3, "ngModel", "ngModelChange"], ["type", "text", "name", "passkey", "inputmode", "numeric", "pattern", "[0-9]*", "autocomplete", "one-time-code", "id", "newpasskey", "placeholder", "Token from Authenticator", "ngModel", "", "required", "", "minlength", "6", "ngbTooltip", "Minimum 6 numbers", 1, "form-control", 3, "disabled"], ["type", "submit", 1, "btn", "btn-outline-primary"]], template: function UserProfileComponent_Template(rf, ctx) { if (rf & 1) {
        _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵelementStart"](0, "div", 0);
        _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵtemplate"](1, UserProfileComponent_div_1_Template, 3, 0, "div", 1);
        _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵelementEnd"]();
        _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵtemplate"](2, UserProfileComponent_ngb_alert_2_Template, 2, 2, "ngb-alert", 2);
        _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵtemplate"](3, UserProfileComponent_form_3_Template, 29, 6, "form", 3);
        _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵtemplate"](4, UserProfileComponent_div_4_Template, 5, 0, "div", 4);
        _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵtemplate"](5, UserProfileComponent_form_5_Template, 25, 8, "form", 3);
    } if (rf & 2) {
        _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵadvance"](1);
        _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵproperty"]("ngIf", ctx.IsLoading);
        _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵadvance"](1);
        _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵproperty"]("ngIf", ctx.ShowMessage);
        _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵadvance"](1);
        _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵproperty"]("ngIf", ctx.UserToEdit);
        _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵadvance"](1);
        _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵproperty"]("ngIf", ctx.UserToEdit && ctx.TwoFactorEnabled);
        _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵadvance"](1);
        _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵproperty"]("ngIf", ctx.UserToEdit && !ctx.TwoFactorEnabled);
    } }, directives: [_angular_common__WEBPACK_IMPORTED_MODULE_5__.NgIf, _ng_bootstrap_ng_bootstrap__WEBPACK_IMPORTED_MODULE_6__.NgbAlert, _angular_forms__WEBPACK_IMPORTED_MODULE_7__["ɵNgNoValidate"], _angular_forms__WEBPACK_IMPORTED_MODULE_7__.NgControlStatusGroup, _angular_forms__WEBPACK_IMPORTED_MODULE_7__.NgForm, _angular_forms__WEBPACK_IMPORTED_MODULE_7__.DefaultValueAccessor, _angular_forms__WEBPACK_IMPORTED_MODULE_7__.RequiredValidator, _angular_forms__WEBPACK_IMPORTED_MODULE_7__.NgControlStatus, _angular_forms__WEBPACK_IMPORTED_MODULE_7__.NgModel, _angular_forms__WEBPACK_IMPORTED_MODULE_7__.EmailValidator, _angular_forms__WEBPACK_IMPORTED_MODULE_7__.PatternValidator, _angular_forms__WEBPACK_IMPORTED_MODULE_7__.CheckboxControlValueAccessor, _angular_router__WEBPACK_IMPORTED_MODULE_4__.RouterLink, _angular_forms__WEBPACK_IMPORTED_MODULE_7__.CheckboxRequiredValidator, _angular_forms__WEBPACK_IMPORTED_MODULE_7__.MinLengthValidator, _ng_bootstrap_ng_bootstrap__WEBPACK_IMPORTED_MODULE_6__.NgbTooltip], pipes: [_angular_common__WEBPACK_IMPORTED_MODULE_5__.AsyncPipe, _shared_secure_image_pipe_pipe__WEBPACK_IMPORTED_MODULE_2__.SecureImagePipe], styles: ["input.ng-invalid.ng-touched[_ngcontent-%COMP%] {\n  border-color: rgba(255, 10, 29, 1) !important;\n  box-shadow: none !important;\n}\n\ninput.ng-invalid.ng-touched[_ngcontent-%COMP%]:focus {\n  box-shadow: 0 0 0 0.2rem rgba(255, 10, 29, 0.25) !important;\n}\n\ninput.ng-valid.ng-touched[_ngcontent-%COMP%] {\n  border-color: rgb(103, 192, 123) !important;\n  box-shadow: none !important;\n}\n\ninput.ng-valid.ng-touched[_ngcontent-%COMP%]:focus {\n  box-shadow: 0 0 0 0.2rem rgba(52, 179, 13, 0.25) !important;\n}\n\nselect.ng-invalid.ng-touched[_ngcontent-%COMP%] {\n  border-color: rgba(255, 10, 29, 1) !important;\n  box-shadow: none !important;\n}\n\nselect.ng-invalid.ng-touched[_ngcontent-%COMP%]:focus {\n  box-shadow: 0 0 0 0.2rem rgba(255, 10, 29, 0.25) !important;\n}\n\nselect.ng-valid.ng-touched[_ngcontent-%COMP%] {\n  border-color: rgb(103, 192, 123) !important;\n  box-shadow: none !important;\n}\n\nselect.ng-valid.ng-touched[_ngcontent-%COMP%]:focus {\n  box-shadow: 0 0 0 0.2rem rgba(52, 179, 13, 0.25) !important;\n}\n\n.image-placeholder[_ngcontent-%COMP%] {\n  background-color: #eee;\n  display: flex;\n  height: 333px;\n  margin: 5px;\n  border-radius: 3px;\n  max-width: 592px;\n}\n\n.image-placeholder[_ngcontent-%COMP%]    > h4[_ngcontent-%COMP%] {\n  align-self: center;\n  text-align: center;\n  width: 100%;\n}\n/*# sourceMappingURL=data:application/json;base64,eyJ2ZXJzaW9uIjozLCJzb3VyY2VzIjpbInVzZXItcHJvZmlsZS5jb21wb25lbnQuY3NzIl0sIm5hbWVzIjpbXSwibWFwcGluZ3MiOiJBQUFBO0VBQ0UsNkNBQTZDO0VBQzdDLDJCQUEyQjtBQUM3Qjs7QUFFQTtFQUNFLDJEQUEyRDtBQUM3RDs7QUFFQTtFQUNFLDJDQUEyQztFQUMzQywyQkFBMkI7QUFDN0I7O0FBRUE7RUFDRSwyREFBMkQ7QUFDN0Q7O0FBRUE7RUFDRSw2Q0FBNkM7RUFDN0MsMkJBQTJCO0FBQzdCOztBQUVBO0VBQ0UsMkRBQTJEO0FBQzdEOztBQUVBO0VBQ0UsMkNBQTJDO0VBQzNDLDJCQUEyQjtBQUM3Qjs7QUFFQTtFQUNFLDJEQUEyRDtBQUM3RDs7QUFFQTtFQUNFLHNCQUFzQjtFQUN0QixhQUFhO0VBQ2IsYUFBYTtFQUNiLFdBQVc7RUFDWCxrQkFBa0I7RUFDbEIsZ0JBQWdCO0FBQ2xCOztBQUVBO0VBQ0Usa0JBQWtCO0VBQ2xCLGtCQUFrQjtFQUNsQixXQUFXO0FBQ2IiLCJmaWxlIjoidXNlci1wcm9maWxlLmNvbXBvbmVudC5jc3MiLCJzb3VyY2VzQ29udGVudCI6WyJpbnB1dC5uZy1pbnZhbGlkLm5nLXRvdWNoZWQge1xuICBib3JkZXItY29sb3I6IHJnYmEoMjU1LCAxMCwgMjksIDEpICFpbXBvcnRhbnQ7XG4gIGJveC1zaGFkb3c6IG5vbmUgIWltcG9ydGFudDtcbn1cblxuaW5wdXQubmctaW52YWxpZC5uZy10b3VjaGVkOmZvY3VzIHtcbiAgYm94LXNoYWRvdzogMCAwIDAgMC4ycmVtIHJnYmEoMjU1LCAxMCwgMjksIDAuMjUpICFpbXBvcnRhbnQ7XG59XG5cbmlucHV0Lm5nLXZhbGlkLm5nLXRvdWNoZWQge1xuICBib3JkZXItY29sb3I6IHJnYigxMDMsIDE5MiwgMTIzKSAhaW1wb3J0YW50O1xuICBib3gtc2hhZG93OiBub25lICFpbXBvcnRhbnQ7XG59XG5cbmlucHV0Lm5nLXZhbGlkLm5nLXRvdWNoZWQ6Zm9jdXMge1xuICBib3gtc2hhZG93OiAwIDAgMCAwLjJyZW0gcmdiYSg1MiwgMTc5LCAxMywgMC4yNSkgIWltcG9ydGFudDtcbn1cblxuc2VsZWN0Lm5nLWludmFsaWQubmctdG91Y2hlZCB7XG4gIGJvcmRlci1jb2xvcjogcmdiYSgyNTUsIDEwLCAyOSwgMSkgIWltcG9ydGFudDtcbiAgYm94LXNoYWRvdzogbm9uZSAhaW1wb3J0YW50O1xufVxuXG5zZWxlY3QubmctaW52YWxpZC5uZy10b3VjaGVkOmZvY3VzIHtcbiAgYm94LXNoYWRvdzogMCAwIDAgMC4ycmVtIHJnYmEoMjU1LCAxMCwgMjksIDAuMjUpICFpbXBvcnRhbnQ7XG59XG5cbnNlbGVjdC5uZy12YWxpZC5uZy10b3VjaGVkIHtcbiAgYm9yZGVyLWNvbG9yOiByZ2IoMTAzLCAxOTIsIDEyMykgIWltcG9ydGFudDtcbiAgYm94LXNoYWRvdzogbm9uZSAhaW1wb3J0YW50O1xufVxuXG5zZWxlY3QubmctdmFsaWQubmctdG91Y2hlZDpmb2N1cyB7XG4gIGJveC1zaGFkb3c6IDAgMCAwIDAuMnJlbSByZ2JhKDUyLCAxNzksIDEzLCAwLjI1KSAhaW1wb3J0YW50O1xufVxuXG4uaW1hZ2UtcGxhY2Vob2xkZXIge1xuICBiYWNrZ3JvdW5kLWNvbG9yOiAjZWVlO1xuICBkaXNwbGF5OiBmbGV4O1xuICBoZWlnaHQ6IDMzM3B4O1xuICBtYXJnaW46IDVweDtcbiAgYm9yZGVyLXJhZGl1czogM3B4O1xuICBtYXgtd2lkdGg6IDU5MnB4O1xufVxuXG4uaW1hZ2UtcGxhY2Vob2xkZXIgPiBoNCB7XG4gIGFsaWduLXNlbGY6IGNlbnRlcjtcbiAgdGV4dC1hbGlnbjogY2VudGVyO1xuICB3aWR0aDogMTAwJTtcbn1cblxuIl19 */"] });


/***/ }),

/***/ 7582:
/*!*****************************************************!*\
  !*** ./src/app/user-profile/user-profile.module.ts ***!
  \*****************************************************/
/***/ ((__unused_webpack_module, __webpack_exports__, __webpack_require__) => {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   "UserProfileModule": () => (/* binding */ UserProfileModule)
/* harmony export */ });
/* harmony import */ var _angular_common__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! @angular/common */ 6362);
/* harmony import */ var _user_profile_component__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ./user-profile.component */ 7960);
/* harmony import */ var _angular_router__WEBPACK_IMPORTED_MODULE_6__ = __webpack_require__(/*! @angular/router */ 2816);
/* harmony import */ var _angular_forms__WEBPACK_IMPORTED_MODULE_4__ = __webpack_require__(/*! @angular/forms */ 587);
/* harmony import */ var _ng_bootstrap_ng_bootstrap__WEBPACK_IMPORTED_MODULE_5__ = __webpack_require__(/*! @ng-bootstrap/ng-bootstrap */ 7544);
/* harmony import */ var _shared_secure_image_pipe_pipe__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! ../shared/secure-image-pipe.pipe */ 9488);
/* harmony import */ var _angular_core__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! @angular/core */ 3184);








const routes = [
    { path: '', component: _user_profile_component__WEBPACK_IMPORTED_MODULE_0__.UserProfileComponent, },
];
class UserProfileModule {
}
UserProfileModule.ɵfac = function UserProfileModule_Factory(t) { return new (t || UserProfileModule)(); };
UserProfileModule.ɵmod = /*@__PURE__*/ _angular_core__WEBPACK_IMPORTED_MODULE_2__["ɵɵdefineNgModule"]({ type: UserProfileModule });
UserProfileModule.ɵinj = /*@__PURE__*/ _angular_core__WEBPACK_IMPORTED_MODULE_2__["ɵɵdefineInjector"]({ imports: [[
            _angular_common__WEBPACK_IMPORTED_MODULE_3__.CommonModule,
            _angular_forms__WEBPACK_IMPORTED_MODULE_4__.FormsModule,
            _ng_bootstrap_ng_bootstrap__WEBPACK_IMPORTED_MODULE_5__.NgbAlertModule,
            _angular_router__WEBPACK_IMPORTED_MODULE_6__.RouterModule.forChild(routes),
            _ng_bootstrap_ng_bootstrap__WEBPACK_IMPORTED_MODULE_5__.NgbTooltipModule
        ]] });
(function () { (typeof ngJitMode === "undefined" || ngJitMode) && _angular_core__WEBPACK_IMPORTED_MODULE_2__["ɵɵsetNgModuleScope"](UserProfileModule, { declarations: [_user_profile_component__WEBPACK_IMPORTED_MODULE_0__.UserProfileComponent,
        _shared_secure_image_pipe_pipe__WEBPACK_IMPORTED_MODULE_1__.SecureImagePipe], imports: [_angular_common__WEBPACK_IMPORTED_MODULE_3__.CommonModule,
        _angular_forms__WEBPACK_IMPORTED_MODULE_4__.FormsModule,
        _ng_bootstrap_ng_bootstrap__WEBPACK_IMPORTED_MODULE_5__.NgbAlertModule, _angular_router__WEBPACK_IMPORTED_MODULE_6__.RouterModule, _ng_bootstrap_ng_bootstrap__WEBPACK_IMPORTED_MODULE_5__.NgbTooltipModule] }); })();


/***/ })

}]);
//# sourceMappingURL=src_app_user-profile_user-profile_module_ts.js.map