"use strict";
(self["webpackChunkshopping_lists_and_recipes"] = self["webpackChunkshopping_lists_and_recipes"] || []).push([["src_app_confirm-email_confirm-email_module_ts"],{

/***/ 1274:
/*!**********************************************************!*\
  !*** ./src/app/confirm-email/confirm-email.component.ts ***!
  \**********************************************************/
/***/ ((__unused_webpack_module, __webpack_exports__, __webpack_require__) => {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   "ConfirmEmailComponent": () => (/* binding */ ConfirmEmailComponent)
/* harmony export */ });
/* harmony import */ var _angular_core__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! @angular/core */ 3184);
/* harmony import */ var _shared_data_storage_service__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ../shared/data-storage.service */ 3649);
/* harmony import */ var _angular_router__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! @angular/router */ 2816);
/* harmony import */ var _angular_common__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! @angular/common */ 6362);
/* harmony import */ var _angular_forms__WEBPACK_IMPORTED_MODULE_4__ = __webpack_require__(/*! @angular/forms */ 587);
/* harmony import */ var _ng_bootstrap_ng_bootstrap__WEBPACK_IMPORTED_MODULE_5__ = __webpack_require__(/*! @ng-bootstrap/ng-bootstrap */ 7544);






function ConfirmEmailComponent_div_1_div_4_input_4_Template(rf, ctx) { if (rf & 1) {
    _angular_core__WEBPACK_IMPORTED_MODULE_1__["ɵɵelement"](0, "input", 18);
} }
function ConfirmEmailComponent_div_1_div_4_input_5_Template(rf, ctx) { if (rf & 1) {
    _angular_core__WEBPACK_IMPORTED_MODULE_1__["ɵɵelement"](0, "input", 19);
} }
function ConfirmEmailComponent_div_1_div_4_div_6_Template(rf, ctx) { if (rf & 1) {
    _angular_core__WEBPACK_IMPORTED_MODULE_1__["ɵɵelementStart"](0, "div", 20)(1, "div", 21)(2, "button", 22);
    _angular_core__WEBPACK_IMPORTED_MODULE_1__["ɵɵtext"](3);
    _angular_core__WEBPACK_IMPORTED_MODULE_1__["ɵɵelementEnd"]()()();
} if (rf & 2) {
    _angular_core__WEBPACK_IMPORTED_MODULE_1__["ɵɵnextContext"]();
    const _r4 = _angular_core__WEBPACK_IMPORTED_MODULE_1__["ɵɵreference"](2);
    const ctx_r7 = _angular_core__WEBPACK_IMPORTED_MODULE_1__["ɵɵnextContext"](2);
    _angular_core__WEBPACK_IMPORTED_MODULE_1__["ɵɵadvance"](2);
    _angular_core__WEBPACK_IMPORTED_MODULE_1__["ɵɵproperty"]("disabled", _r4.invalid);
    _angular_core__WEBPACK_IMPORTED_MODULE_1__["ɵɵadvance"](1);
    _angular_core__WEBPACK_IMPORTED_MODULE_1__["ɵɵtextInterpolate"](ctx_r7.ResetPasswordMode && ctx_r7.Token ? "Save" : "Send");
} }
function ConfirmEmailComponent_div_1_div_4_Template(rf, ctx) { if (rf & 1) {
    const _r9 = _angular_core__WEBPACK_IMPORTED_MODULE_1__["ɵɵgetCurrentView"]();
    _angular_core__WEBPACK_IMPORTED_MODULE_1__["ɵɵelementStart"](0, "div", 11)(1, "form", 12, 13);
    _angular_core__WEBPACK_IMPORTED_MODULE_1__["ɵɵlistener"]("ngSubmit", function ConfirmEmailComponent_div_1_div_4_Template_form_ngSubmit_1_listener() { _angular_core__WEBPACK_IMPORTED_MODULE_1__["ɵɵrestoreView"](_r9); const _r4 = _angular_core__WEBPACK_IMPORTED_MODULE_1__["ɵɵreference"](2); const ctx_r8 = _angular_core__WEBPACK_IMPORTED_MODULE_1__["ɵɵnextContext"](2); return ctx_r8.OnSubmitForm(_r4); });
    _angular_core__WEBPACK_IMPORTED_MODULE_1__["ɵɵelementStart"](3, "div", 14);
    _angular_core__WEBPACK_IMPORTED_MODULE_1__["ɵɵtemplate"](4, ConfirmEmailComponent_div_1_div_4_input_4_Template, 1, 0, "input", 15);
    _angular_core__WEBPACK_IMPORTED_MODULE_1__["ɵɵtemplate"](5, ConfirmEmailComponent_div_1_div_4_input_5_Template, 1, 0, "input", 16);
    _angular_core__WEBPACK_IMPORTED_MODULE_1__["ɵɵelementEnd"]();
    _angular_core__WEBPACK_IMPORTED_MODULE_1__["ɵɵtemplate"](6, ConfirmEmailComponent_div_1_div_4_div_6_Template, 4, 2, "div", 17);
    _angular_core__WEBPACK_IMPORTED_MODULE_1__["ɵɵelementEnd"]()();
} if (rf & 2) {
    const ctx_r2 = _angular_core__WEBPACK_IMPORTED_MODULE_1__["ɵɵnextContext"](2);
    _angular_core__WEBPACK_IMPORTED_MODULE_1__["ɵɵadvance"](4);
    _angular_core__WEBPACK_IMPORTED_MODULE_1__["ɵɵproperty"]("ngIf", !ctx_r2.Token);
    _angular_core__WEBPACK_IMPORTED_MODULE_1__["ɵɵadvance"](1);
    _angular_core__WEBPACK_IMPORTED_MODULE_1__["ɵɵproperty"]("ngIf", ctx_r2.ResetPasswordMode && ctx_r2.Token);
    _angular_core__WEBPACK_IMPORTED_MODULE_1__["ɵɵadvance"](1);
    _angular_core__WEBPACK_IMPORTED_MODULE_1__["ɵɵproperty"]("ngIf", !ctx_r2.Token || ctx_r2.ResetPasswordMode);
} }
function ConfirmEmailComponent_div_1_div_9_ngb_alert_1_Template(rf, ctx) { if (rf & 1) {
    const _r12 = _angular_core__WEBPACK_IMPORTED_MODULE_1__["ɵɵgetCurrentView"]();
    _angular_core__WEBPACK_IMPORTED_MODULE_1__["ɵɵelementStart"](0, "ngb-alert", 25);
    _angular_core__WEBPACK_IMPORTED_MODULE_1__["ɵɵlistener"]("close", function ConfirmEmailComponent_div_1_div_9_ngb_alert_1_Template_ngb_alert_close_0_listener() { _angular_core__WEBPACK_IMPORTED_MODULE_1__["ɵɵrestoreView"](_r12); const ctx_r11 = _angular_core__WEBPACK_IMPORTED_MODULE_1__["ɵɵnextContext"](3); return ctx_r11.ShowMessage = false; });
    _angular_core__WEBPACK_IMPORTED_MODULE_1__["ɵɵtext"](1);
    _angular_core__WEBPACK_IMPORTED_MODULE_1__["ɵɵelementEnd"]();
} if (rf & 2) {
    const ctx_r10 = _angular_core__WEBPACK_IMPORTED_MODULE_1__["ɵɵnextContext"](3);
    _angular_core__WEBPACK_IMPORTED_MODULE_1__["ɵɵproperty"]("type", ctx_r10.MessageType);
    _angular_core__WEBPACK_IMPORTED_MODULE_1__["ɵɵadvance"](1);
    _angular_core__WEBPACK_IMPORTED_MODULE_1__["ɵɵtextInterpolate"](ctx_r10.ResponseFromBackend.Error.Message);
} }
function ConfirmEmailComponent_div_1_div_9_Template(rf, ctx) { if (rf & 1) {
    _angular_core__WEBPACK_IMPORTED_MODULE_1__["ɵɵelementStart"](0, "div", 23);
    _angular_core__WEBPACK_IMPORTED_MODULE_1__["ɵɵtemplate"](1, ConfirmEmailComponent_div_1_div_9_ngb_alert_1_Template, 2, 2, "ngb-alert", 24);
    _angular_core__WEBPACK_IMPORTED_MODULE_1__["ɵɵelementEnd"]();
} if (rf & 2) {
    const ctx_r3 = _angular_core__WEBPACK_IMPORTED_MODULE_1__["ɵɵnextContext"](2);
    _angular_core__WEBPACK_IMPORTED_MODULE_1__["ɵɵadvance"](1);
    _angular_core__WEBPACK_IMPORTED_MODULE_1__["ɵɵproperty"]("ngIf", ctx_r3.ShowMessage);
} }
function ConfirmEmailComponent_div_1_Template(rf, ctx) { if (rf & 1) {
    _angular_core__WEBPACK_IMPORTED_MODULE_1__["ɵɵelementStart"](0, "div", 4)(1, "div", 5)(2, "h3");
    _angular_core__WEBPACK_IMPORTED_MODULE_1__["ɵɵtext"](3);
    _angular_core__WEBPACK_IMPORTED_MODULE_1__["ɵɵelementEnd"]()();
    _angular_core__WEBPACK_IMPORTED_MODULE_1__["ɵɵtemplate"](4, ConfirmEmailComponent_div_1_div_4_Template, 7, 3, "div", 6);
    _angular_core__WEBPACK_IMPORTED_MODULE_1__["ɵɵelementStart"](5, "div", 7)(6, "div", 8)(7, "a", 9);
    _angular_core__WEBPACK_IMPORTED_MODULE_1__["ɵɵtext"](8, "Back");
    _angular_core__WEBPACK_IMPORTED_MODULE_1__["ɵɵelementEnd"]()()();
    _angular_core__WEBPACK_IMPORTED_MODULE_1__["ɵɵtemplate"](9, ConfirmEmailComponent_div_1_div_9_Template, 2, 1, "div", 10);
    _angular_core__WEBPACK_IMPORTED_MODULE_1__["ɵɵelementEnd"]();
} if (rf & 2) {
    const ctx_r0 = _angular_core__WEBPACK_IMPORTED_MODULE_1__["ɵɵnextContext"]();
    _angular_core__WEBPACK_IMPORTED_MODULE_1__["ɵɵadvance"](3);
    _angular_core__WEBPACK_IMPORTED_MODULE_1__["ɵɵtextInterpolate"](!ctx_r0.ResetPasswordMode ? "Confirm email" : "Reset password");
    _angular_core__WEBPACK_IMPORTED_MODULE_1__["ɵɵadvance"](1);
    _angular_core__WEBPACK_IMPORTED_MODULE_1__["ɵɵproperty"]("ngIf", !ctx_r0.Token || ctx_r0.ResetPasswordMode);
    _angular_core__WEBPACK_IMPORTED_MODULE_1__["ɵɵadvance"](5);
    _angular_core__WEBPACK_IMPORTED_MODULE_1__["ɵɵproperty"]("ngIf", ctx_r0.ShowMessage);
} }
function ConfirmEmailComponent_div_3_Template(rf, ctx) { if (rf & 1) {
    _angular_core__WEBPACK_IMPORTED_MODULE_1__["ɵɵelementStart"](0, "div", 26)(1, "span", 27);
    _angular_core__WEBPACK_IMPORTED_MODULE_1__["ɵɵtext"](2, "Loading...");
    _angular_core__WEBPACK_IMPORTED_MODULE_1__["ɵɵelementEnd"]()();
} }
class ConfirmEmailComponent {
    constructor(DataServ, activeroute, router) {
        this.DataServ = DataServ;
        this.activeroute = activeroute;
        this.router = router;
        this.IsLoading = false;
    }
    ngOnDestroy() {
        this.RecivedErrorSub.unsubscribe();
        this.RecivedResponseSub.unsubscribe();
        this.DataServiceSub.unsubscribe();
    }
    ngOnInit() {
        this.DataServiceSub = this.DataServ.LoadingData.subscribe((State) => {
            this.IsLoading = State;
        });
        this.RecivedErrorSub = this.DataServ.RecivedError.subscribe((response) => {
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
        this.activeroute.queryParams.subscribe((Qparams) => {
            this.Token = Qparams.Token;
            const cururl = this.getUrlWithoutParams();
            this.ResetPasswordMode = cururl === '/reset-password';
            if (!this.ResetPasswordMode) {
                if (this.Token) {
                    this.DataServ.ConfirmEmail(this.Token);
                }
            }
        });
    }
    getUrlWithoutParams() {
        const urlTree = this.router.parseUrl(this.router.url);
        urlTree.queryParams = {};
        return urlTree.toString();
    }
    OnSubmitForm(ResendConfEmailForm) {
        if (ResendConfEmailForm.valid) {
            this.IsLoading = true;
            if (this.ResetPasswordMode && this.Token) {
                this.DataServ.SubmitNewPassword(this.Token, ResendConfEmailForm.value.newpassword);
            }
            else {
                if (this.ResetPasswordMode) {
                    this.DataServ.SendEmailResetPassword(ResendConfEmailForm.value.email);
                }
                else {
                    this.DataServ.SendEmailConfirmEmail(ResendConfEmailForm.value.email);
                }
            }
            ResendConfEmailForm.reset();
        }
    }
}
ConfirmEmailComponent.ɵfac = function ConfirmEmailComponent_Factory(t) { return new (t || ConfirmEmailComponent)(_angular_core__WEBPACK_IMPORTED_MODULE_1__["ɵɵdirectiveInject"](_shared_data_storage_service__WEBPACK_IMPORTED_MODULE_0__.DataStorageService), _angular_core__WEBPACK_IMPORTED_MODULE_1__["ɵɵdirectiveInject"](_angular_router__WEBPACK_IMPORTED_MODULE_2__.ActivatedRoute), _angular_core__WEBPACK_IMPORTED_MODULE_1__["ɵɵdirectiveInject"](_angular_router__WEBPACK_IMPORTED_MODULE_2__.Router)); };
ConfirmEmailComponent.ɵcmp = /*@__PURE__*/ _angular_core__WEBPACK_IMPORTED_MODULE_1__["ɵɵdefineComponent"]({ type: ConfirmEmailComponent, selectors: [["app-confirm-email"]], decls: 4, vars: 2, consts: [[1, "main-parent-login"], ["class", "card mx-auto", "style", "margin-top: 3px;", 4, "ngIf"], [1, "text-center"], ["class", "spinner-border", "role", "status", 4, "ngIf"], [1, "card", "mx-auto", 2, "margin-top", "3px"], [1, "card-header"], ["class", "card-body", 4, "ngIf"], [1, "card-footer"], [1, "d-flex", "justify-content-center"], ["href", "/"], ["style", "padding: 3px", 4, "ngIf"], [1, "card-body"], [3, "ngSubmit"], ["ResendConfEmailForm", "ngForm"], [1, "input-group", "form-group"], ["type", "email", "name", "email", "class", "form-control", "placeholder", "E-mail", "placement", "right", "ngbTooltip", "Registered email required", "ngModel", "", "email", "", "required", "", 4, "ngIf"], ["type", "password", "name", "newpassword", "class", "form-control", "placeholder", "New password", "ngModel", "", "minlength", "6", "placement", "right", "ngbTooltip", "Minimum 6 chars", "required", "", 4, "ngIf"], ["class", "form-group float-right", 4, "ngIf"], ["type", "email", "name", "email", "placeholder", "E-mail", "placement", "right", "ngbTooltip", "Registered email required", "ngModel", "", "email", "", "required", "", 1, "form-control"], ["type", "password", "name", "newpassword", "placeholder", "New password", "ngModel", "", "minlength", "6", "placement", "right", "ngbTooltip", "Minimum 6 chars", "required", "", 1, "form-control"], [1, "form-group", "float-right"], [1, "input-group"], ["type", "submit", 1, "btn", "btn", "btn-outline-primary", 3, "disabled"], [2, "padding", "3px"], [3, "type", "close", 4, "ngIf"], [3, "type", "close"], ["role", "status", 1, "spinner-border"], [1, "sr-only"]], template: function ConfirmEmailComponent_Template(rf, ctx) { if (rf & 1) {
        _angular_core__WEBPACK_IMPORTED_MODULE_1__["ɵɵelementStart"](0, "div", 0);
        _angular_core__WEBPACK_IMPORTED_MODULE_1__["ɵɵtemplate"](1, ConfirmEmailComponent_div_1_Template, 10, 3, "div", 1);
        _angular_core__WEBPACK_IMPORTED_MODULE_1__["ɵɵelementStart"](2, "div", 2);
        _angular_core__WEBPACK_IMPORTED_MODULE_1__["ɵɵtemplate"](3, ConfirmEmailComponent_div_3_Template, 3, 0, "div", 3);
        _angular_core__WEBPACK_IMPORTED_MODULE_1__["ɵɵelementEnd"]()();
    } if (rf & 2) {
        _angular_core__WEBPACK_IMPORTED_MODULE_1__["ɵɵadvance"](1);
        _angular_core__WEBPACK_IMPORTED_MODULE_1__["ɵɵproperty"]("ngIf", !ctx.IsLoading);
        _angular_core__WEBPACK_IMPORTED_MODULE_1__["ɵɵadvance"](2);
        _angular_core__WEBPACK_IMPORTED_MODULE_1__["ɵɵproperty"]("ngIf", ctx.IsLoading);
    } }, directives: [_angular_common__WEBPACK_IMPORTED_MODULE_3__.NgIf, _angular_forms__WEBPACK_IMPORTED_MODULE_4__["ɵNgNoValidate"], _angular_forms__WEBPACK_IMPORTED_MODULE_4__.NgControlStatusGroup, _angular_forms__WEBPACK_IMPORTED_MODULE_4__.NgForm, _angular_forms__WEBPACK_IMPORTED_MODULE_4__.DefaultValueAccessor, _ng_bootstrap_ng_bootstrap__WEBPACK_IMPORTED_MODULE_5__.NgbTooltip, _angular_forms__WEBPACK_IMPORTED_MODULE_4__.NgControlStatus, _angular_forms__WEBPACK_IMPORTED_MODULE_4__.NgModel, _angular_forms__WEBPACK_IMPORTED_MODULE_4__.EmailValidator, _angular_forms__WEBPACK_IMPORTED_MODULE_4__.RequiredValidator, _angular_forms__WEBPACK_IMPORTED_MODULE_4__.MinLengthValidator, _ng_bootstrap_ng_bootstrap__WEBPACK_IMPORTED_MODULE_5__.NgbAlert], styles: [".main-parent-login[_ngcontent-%COMP%] {\n  position: absolute;\n  top: 50%;\n  left: 50%;\n  transform: translate(-50%, -50%);\n}\n\n.card[_ngcontent-%COMP%] {\n  width: 300px;\n}\n\n.remember[_ngcontent-%COMP%]   input[_ngcontent-%COMP%] {\n  width: 20px;\n  height: 20px;\n  margin-left: 15px;\n  margin-right: 5px;\n}\n\ninput.ng-invalid.ng-touched[_ngcontent-%COMP%] {\n  border-color: rgba(255, 10, 29, 1) !important;\n  box-shadow: none !important;\n}\n\ninput.ng-invalid.ng-touched[_ngcontent-%COMP%]:focus {\n  box-shadow: 0 0 0 0.2rem rgba(255, 10, 29, 0.25) !important;\n}\n\ninput.ng-valid.ng-touched[_ngcontent-%COMP%] {\n  border-color: rgb(103, 192, 123) !important;\n  box-shadow: none !important;\n}\n\ninput.ng-valid.ng-touched[_ngcontent-%COMP%]:focus {\n  box-shadow: 0 0 0 0.2rem rgba(52, 179, 13, 0.25) !important;\n}\n/*# sourceMappingURL=data:application/json;base64,eyJ2ZXJzaW9uIjozLCJzb3VyY2VzIjpbImNvbmZpcm0tZW1haWwuY29tcG9uZW50LmNzcyJdLCJuYW1lcyI6W10sIm1hcHBpbmdzIjoiQUFBQTtFQUNFLGtCQUFrQjtFQUNsQixRQUFRO0VBQ1IsU0FBUztFQUNULGdDQUFnQztBQUNsQzs7QUFFQTtFQUNFLFlBQVk7QUFDZDs7QUFFQTtFQUNFLFdBQVc7RUFDWCxZQUFZO0VBQ1osaUJBQWlCO0VBQ2pCLGlCQUFpQjtBQUNuQjs7QUFFQTtFQUNFLDZDQUE2QztFQUM3QywyQkFBMkI7QUFDN0I7O0FBRUE7RUFDRSwyREFBMkQ7QUFDN0Q7O0FBRUE7RUFDRSwyQ0FBMkM7RUFDM0MsMkJBQTJCO0FBQzdCOztBQUVBO0VBQ0UsMkRBQTJEO0FBQzdEIiwiZmlsZSI6ImNvbmZpcm0tZW1haWwuY29tcG9uZW50LmNzcyIsInNvdXJjZXNDb250ZW50IjpbIi5tYWluLXBhcmVudC1sb2dpbiB7XG4gIHBvc2l0aW9uOiBhYnNvbHV0ZTtcbiAgdG9wOiA1MCU7XG4gIGxlZnQ6IDUwJTtcbiAgdHJhbnNmb3JtOiB0cmFuc2xhdGUoLTUwJSwgLTUwJSk7XG59XG5cbi5jYXJkIHtcbiAgd2lkdGg6IDMwMHB4O1xufVxuXG4ucmVtZW1iZXIgaW5wdXQge1xuICB3aWR0aDogMjBweDtcbiAgaGVpZ2h0OiAyMHB4O1xuICBtYXJnaW4tbGVmdDogMTVweDtcbiAgbWFyZ2luLXJpZ2h0OiA1cHg7XG59XG5cbmlucHV0Lm5nLWludmFsaWQubmctdG91Y2hlZCB7XG4gIGJvcmRlci1jb2xvcjogcmdiYSgyNTUsIDEwLCAyOSwgMSkgIWltcG9ydGFudDtcbiAgYm94LXNoYWRvdzogbm9uZSAhaW1wb3J0YW50O1xufVxuXG5pbnB1dC5uZy1pbnZhbGlkLm5nLXRvdWNoZWQ6Zm9jdXMge1xuICBib3gtc2hhZG93OiAwIDAgMCAwLjJyZW0gcmdiYSgyNTUsIDEwLCAyOSwgMC4yNSkgIWltcG9ydGFudDtcbn1cblxuaW5wdXQubmctdmFsaWQubmctdG91Y2hlZCB7XG4gIGJvcmRlci1jb2xvcjogcmdiKDEwMywgMTkyLCAxMjMpICFpbXBvcnRhbnQ7XG4gIGJveC1zaGFkb3c6IG5vbmUgIWltcG9ydGFudDtcbn1cblxuaW5wdXQubmctdmFsaWQubmctdG91Y2hlZDpmb2N1cyB7XG4gIGJveC1zaGFkb3c6IDAgMCAwIDAuMnJlbSByZ2JhKDUyLCAxNzksIDEzLCAwLjI1KSAhaW1wb3J0YW50O1xufVxuIl19 */"] });


/***/ }),

/***/ 2994:
/*!*******************************************************!*\
  !*** ./src/app/confirm-email/confirm-email.module.ts ***!
  \*******************************************************/
/***/ ((__unused_webpack_module, __webpack_exports__, __webpack_require__) => {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   "ConfirmEmailModule": () => (/* binding */ ConfirmEmailModule)
/* harmony export */ });
/* harmony import */ var _angular_common__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! @angular/common */ 6362);
/* harmony import */ var _angular_router__WEBPACK_IMPORTED_MODULE_5__ = __webpack_require__(/*! @angular/router */ 2816);
/* harmony import */ var _confirm_email_component__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ./confirm-email.component */ 1274);
/* harmony import */ var _angular_forms__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! @angular/forms */ 587);
/* harmony import */ var _ng_bootstrap_ng_bootstrap__WEBPACK_IMPORTED_MODULE_4__ = __webpack_require__(/*! @ng-bootstrap/ng-bootstrap */ 7544);
/* harmony import */ var _angular_core__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! @angular/core */ 3184);







const routes = [
    { path: '', component: _confirm_email_component__WEBPACK_IMPORTED_MODULE_0__.ConfirmEmailComponent, },
];
class ConfirmEmailModule {
}
ConfirmEmailModule.ɵfac = function ConfirmEmailModule_Factory(t) { return new (t || ConfirmEmailModule)(); };
ConfirmEmailModule.ɵmod = /*@__PURE__*/ _angular_core__WEBPACK_IMPORTED_MODULE_1__["ɵɵdefineNgModule"]({ type: ConfirmEmailModule });
ConfirmEmailModule.ɵinj = /*@__PURE__*/ _angular_core__WEBPACK_IMPORTED_MODULE_1__["ɵɵdefineInjector"]({ imports: [[
            _angular_common__WEBPACK_IMPORTED_MODULE_2__.CommonModule,
            _angular_forms__WEBPACK_IMPORTED_MODULE_3__.FormsModule,
            _ng_bootstrap_ng_bootstrap__WEBPACK_IMPORTED_MODULE_4__.NgbAlertModule,
            _angular_router__WEBPACK_IMPORTED_MODULE_5__.RouterModule.forChild(routes),
            _ng_bootstrap_ng_bootstrap__WEBPACK_IMPORTED_MODULE_4__.NgbTooltipModule
        ]] });
(function () { (typeof ngJitMode === "undefined" || ngJitMode) && _angular_core__WEBPACK_IMPORTED_MODULE_1__["ɵɵsetNgModuleScope"](ConfirmEmailModule, { declarations: [_confirm_email_component__WEBPACK_IMPORTED_MODULE_0__.ConfirmEmailComponent], imports: [_angular_common__WEBPACK_IMPORTED_MODULE_2__.CommonModule,
        _angular_forms__WEBPACK_IMPORTED_MODULE_3__.FormsModule,
        _ng_bootstrap_ng_bootstrap__WEBPACK_IMPORTED_MODULE_4__.NgbAlertModule, _angular_router__WEBPACK_IMPORTED_MODULE_5__.RouterModule, _ng_bootstrap_ng_bootstrap__WEBPACK_IMPORTED_MODULE_4__.NgbTooltipModule] }); })();


/***/ })

}]);
//# sourceMappingURL=src_app_confirm-email_confirm-email_module_ts.js.map