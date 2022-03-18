"use strict";
(self["webpackChunkshopping_lists_and_recipes"] = self["webpackChunkshopping_lists_and_recipes"] || []).push([["main"],{

/***/ 158:
/*!***************************************!*\
  !*** ./src/app/app-routing.module.ts ***!
  \***************************************/
/***/ ((__unused_webpack_module, __webpack_exports__, __webpack_require__) => {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   "AppRoutingModule": () => (/* binding */ AppRoutingModule)
/* harmony export */ });
/* harmony import */ var _angular_router__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! @angular/router */ 2816);
/* harmony import */ var _angular_core__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! @angular/core */ 3184);



const routes = [
    { path: '', redirectTo: '/recipes/1', pathMatch: 'full' },
    { path: 'recipes', loadChildren: () => Promise.all(/*! import() */[__webpack_require__.e("default-src_app_shared_data-storage_service_ts"), __webpack_require__.e("common"), __webpack_require__.e("src_app_recipes_recipes_module_ts")]).then(__webpack_require__.bind(__webpack_require__, /*! ./recipes/recipes.module */ 7776)).then(m => m.RecipesModule) },
    { path: 'shopping-list', loadChildren: () => Promise.all(/*! import() */[__webpack_require__.e("default-src_app_shared_data-storage_service_ts"), __webpack_require__.e("common"), __webpack_require__.e("src_app_shopping-list_shopping-list_module_ts")]).then(__webpack_require__.bind(__webpack_require__, /*! ./shopping-list/shopping-list.module */ 6673)).then(m => m.ShoppingListModule) },
    { path: 'auth', loadChildren: () => __webpack_require__.e(/*! import() */ "src_app_auth_auth-feature_module_ts").then(__webpack_require__.bind(__webpack_require__, /*! ./auth/auth-feature.module */ 2773)).then(m => m.AuthFeatureModule) },
    { path: 'admin', loadChildren: () => Promise.all(/*! import() */[__webpack_require__.e("default-src_app_shared_data-storage_service_ts"), __webpack_require__.e("common"), __webpack_require__.e("src_app_admin_admin_module_ts")]).then(__webpack_require__.bind(__webpack_require__, /*! ./admin/admin.module */ 7095)).then(m => m.AdminModule) },
    { path: 'confirm-email', loadChildren: () => Promise.all(/*! import() */[__webpack_require__.e("default-src_app_shared_data-storage_service_ts"), __webpack_require__.e("src_app_confirm-email_confirm-email_module_ts")]).then(__webpack_require__.bind(__webpack_require__, /*! ./confirm-email/confirm-email.module */ 2994)).then(m => m.ConfirmEmailModule) },
    { path: 'reset-password', loadChildren: () => Promise.all(/*! import() */[__webpack_require__.e("default-src_app_shared_data-storage_service_ts"), __webpack_require__.e("src_app_confirm-email_confirm-email_module_ts")]).then(__webpack_require__.bind(__webpack_require__, /*! ./confirm-email/confirm-email.module */ 2994)).then(m => m.ConfirmEmailModule) },
    { path: 'profile', loadChildren: () => Promise.all(/*! import() */[__webpack_require__.e("default-src_app_shared_data-storage_service_ts"), __webpack_require__.e("src_app_user-profile_user-profile_module_ts")]).then(__webpack_require__.bind(__webpack_require__, /*! ./user-profile/user-profile.module */ 7582)).then(m => m.UserProfileModule) },
    { path: 'totp', loadChildren: () => __webpack_require__.e(/*! import() */ "src_app_totp_totp_module_ts").then(__webpack_require__.bind(__webpack_require__, /*! ./totp/totp.module */ 9782)).then(m => m.TotpModule) },
];
class AppRoutingModule {
}
AppRoutingModule.ɵfac = function AppRoutingModule_Factory(t) { return new (t || AppRoutingModule)(); };
AppRoutingModule.ɵmod = /*@__PURE__*/ _angular_core__WEBPACK_IMPORTED_MODULE_0__["ɵɵdefineNgModule"]({ type: AppRoutingModule });
AppRoutingModule.ɵinj = /*@__PURE__*/ _angular_core__WEBPACK_IMPORTED_MODULE_0__["ɵɵdefineInjector"]({ imports: [[_angular_router__WEBPACK_IMPORTED_MODULE_1__.RouterModule.forRoot(routes, { preloadingStrategy: _angular_router__WEBPACK_IMPORTED_MODULE_1__.PreloadAllModules, relativeLinkResolution: 'legacy' })], _angular_router__WEBPACK_IMPORTED_MODULE_1__.RouterModule] });
(function () { (typeof ngJitMode === "undefined" || ngJitMode) && _angular_core__WEBPACK_IMPORTED_MODULE_0__["ɵɵsetNgModuleScope"](AppRoutingModule, { imports: [_angular_router__WEBPACK_IMPORTED_MODULE_1__.RouterModule], exports: [_angular_router__WEBPACK_IMPORTED_MODULE_1__.RouterModule] }); })();


/***/ }),

/***/ 5041:
/*!**********************************!*\
  !*** ./src/app/app.component.ts ***!
  \**********************************/
/***/ ((__unused_webpack_module, __webpack_exports__, __webpack_require__) => {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   "AppComponent": () => (/* binding */ AppComponent)
/* harmony export */ });
/* harmony import */ var _angular_core__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! @angular/core */ 3184);
/* harmony import */ var _auth_auth_service__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ./auth/auth.service */ 384);
/* harmony import */ var _header_header_component__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! ./header/header.component */ 3482);
/* harmony import */ var _angular_router__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! @angular/router */ 2816);




class AppComponent {
    constructor(auth) {
        this.auth = auth;
    }
    ngOnInit() {
        this.auth.AutoSignIn();
    }
}
AppComponent.ɵfac = function AppComponent_Factory(t) { return new (t || AppComponent)(_angular_core__WEBPACK_IMPORTED_MODULE_2__["ɵɵdirectiveInject"](_auth_auth_service__WEBPACK_IMPORTED_MODULE_0__.AuthService)); };
AppComponent.ɵcmp = /*@__PURE__*/ _angular_core__WEBPACK_IMPORTED_MODULE_2__["ɵɵdefineComponent"]({ type: AppComponent, selectors: [["app-root"]], decls: 3, vars: 0, consts: [[1, "container-fluid"]], template: function AppComponent_Template(rf, ctx) { if (rf & 1) {
        _angular_core__WEBPACK_IMPORTED_MODULE_2__["ɵɵelementStart"](0, "div", 0);
        _angular_core__WEBPACK_IMPORTED_MODULE_2__["ɵɵelement"](1, "app-header")(2, "router-outlet");
        _angular_core__WEBPACK_IMPORTED_MODULE_2__["ɵɵelementEnd"]();
    } }, directives: [_header_header_component__WEBPACK_IMPORTED_MODULE_1__.HeaderComponent, _angular_router__WEBPACK_IMPORTED_MODULE_3__.RouterOutlet], styles: ["\n/*# sourceMappingURL=data:application/json;base64,eyJ2ZXJzaW9uIjozLCJzb3VyY2VzIjpbXSwibmFtZXMiOltdLCJtYXBwaW5ncyI6IiIsImZpbGUiOiJhcHAuY29tcG9uZW50LmNzcyJ9 */"] });


/***/ }),

/***/ 6747:
/*!*******************************!*\
  !*** ./src/app/app.module.ts ***!
  \*******************************/
/***/ ((__unused_webpack_module, __webpack_exports__, __webpack_require__) => {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   "AppModule": () => (/* binding */ AppModule)
/* harmony export */ });
/* harmony import */ var _angular_platform_browser__WEBPACK_IMPORTED_MODULE_7__ = __webpack_require__(/*! @angular/platform-browser */ 318);
/* harmony import */ var _angular_common_http__WEBPACK_IMPORTED_MODULE_6__ = __webpack_require__(/*! @angular/common/http */ 8784);
/* harmony import */ var _app_component__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ./app.component */ 5041);
/* harmony import */ var _header_header_component__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! ./header/header.component */ 3482);
/* harmony import */ var _ng_bootstrap_ng_bootstrap__WEBPACK_IMPORTED_MODULE_8__ = __webpack_require__(/*! @ng-bootstrap/ng-bootstrap */ 7544);
/* harmony import */ var _app_routing_module__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! ./app-routing.module */ 158);
/* harmony import */ var _auth_auth_interceptor_service__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! ./auth/auth-interceptor.service */ 2458);
/* harmony import */ var _shared_shared_module__WEBPACK_IMPORTED_MODULE_4__ = __webpack_require__(/*! ./shared/shared.module */ 4466);
/* harmony import */ var _angular_forms__WEBPACK_IMPORTED_MODULE_9__ = __webpack_require__(/*! @angular/forms */ 587);
/* harmony import */ var _angular_platform_browser_animations__WEBPACK_IMPORTED_MODULE_10__ = __webpack_require__(/*! @angular/platform-browser/animations */ 3598);
/* harmony import */ var _angular_core__WEBPACK_IMPORTED_MODULE_5__ = __webpack_require__(/*! @angular/core */ 3184);











class AppModule {
}
AppModule.ɵfac = function AppModule_Factory(t) { return new (t || AppModule)(); };
AppModule.ɵmod = /*@__PURE__*/ _angular_core__WEBPACK_IMPORTED_MODULE_5__["ɵɵdefineNgModule"]({ type: AppModule, bootstrap: [_app_component__WEBPACK_IMPORTED_MODULE_0__.AppComponent] });
AppModule.ɵinj = /*@__PURE__*/ _angular_core__WEBPACK_IMPORTED_MODULE_5__["ɵɵdefineInjector"]({ providers: [{ provide: _angular_common_http__WEBPACK_IMPORTED_MODULE_6__.HTTP_INTERCEPTORS, useClass: _auth_auth_interceptor_service__WEBPACK_IMPORTED_MODULE_3__.AuthInterceptorService, multi: true }], imports: [[
            _angular_platform_browser__WEBPACK_IMPORTED_MODULE_7__.BrowserModule,
            _ng_bootstrap_ng_bootstrap__WEBPACK_IMPORTED_MODULE_8__.NgbTooltipModule,
            _ng_bootstrap_ng_bootstrap__WEBPACK_IMPORTED_MODULE_8__.NgbDropdownModule,
            _ng_bootstrap_ng_bootstrap__WEBPACK_IMPORTED_MODULE_8__.NgbPaginationModule,
            _ng_bootstrap_ng_bootstrap__WEBPACK_IMPORTED_MODULE_8__.NgbAlertModule,
            _app_routing_module__WEBPACK_IMPORTED_MODULE_2__.AppRoutingModule,
            _angular_common_http__WEBPACK_IMPORTED_MODULE_6__.HttpClientModule,
            _angular_forms__WEBPACK_IMPORTED_MODULE_9__.FormsModule,
            _shared_shared_module__WEBPACK_IMPORTED_MODULE_4__.SharedModule,
            _angular_platform_browser_animations__WEBPACK_IMPORTED_MODULE_10__.BrowserAnimationsModule
        ]] });
(function () { (typeof ngJitMode === "undefined" || ngJitMode) && _angular_core__WEBPACK_IMPORTED_MODULE_5__["ɵɵsetNgModuleScope"](AppModule, { declarations: [_app_component__WEBPACK_IMPORTED_MODULE_0__.AppComponent,
        _header_header_component__WEBPACK_IMPORTED_MODULE_1__.HeaderComponent], imports: [_angular_platform_browser__WEBPACK_IMPORTED_MODULE_7__.BrowserModule,
        _ng_bootstrap_ng_bootstrap__WEBPACK_IMPORTED_MODULE_8__.NgbTooltipModule,
        _ng_bootstrap_ng_bootstrap__WEBPACK_IMPORTED_MODULE_8__.NgbDropdownModule,
        _ng_bootstrap_ng_bootstrap__WEBPACK_IMPORTED_MODULE_8__.NgbPaginationModule,
        _ng_bootstrap_ng_bootstrap__WEBPACK_IMPORTED_MODULE_8__.NgbAlertModule,
        _app_routing_module__WEBPACK_IMPORTED_MODULE_2__.AppRoutingModule,
        _angular_common_http__WEBPACK_IMPORTED_MODULE_6__.HttpClientModule,
        _angular_forms__WEBPACK_IMPORTED_MODULE_9__.FormsModule,
        _shared_shared_module__WEBPACK_IMPORTED_MODULE_4__.SharedModule,
        _angular_platform_browser_animations__WEBPACK_IMPORTED_MODULE_10__.BrowserAnimationsModule] }); })();


/***/ }),

/***/ 2458:
/*!**************************************************!*\
  !*** ./src/app/auth/auth-interceptor.service.ts ***!
  \**************************************************/
/***/ ((__unused_webpack_module, __webpack_exports__, __webpack_require__) => {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   "AuthInterceptorService": () => (/* binding */ AuthInterceptorService)
/* harmony export */ });
/* harmony import */ var src_environments_environment__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! src/environments/environment */ 2340);
/* harmony import */ var _angular_core__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! @angular/core */ 3184);
/* harmony import */ var _auth_service__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! ./auth.service */ 384);



class AuthInterceptorService {
    constructor(auth) {
        this.auth = auth;
    }
    intercept(req, next) {
        if (this.auth.CheckFirstFactorPassed()) {
            const modreq = req.clone({ headers: req.headers.set('Auth', this.auth.GetUserToken())
                    .set('ApiKey', src_environments_environment__WEBPACK_IMPORTED_MODULE_0__.environment.ApiKey) });
            return next.handle(modreq);
        }
        else {
            const modreq = req.clone({ headers: req.headers.set('ApiKey', src_environments_environment__WEBPACK_IMPORTED_MODULE_0__.environment.ApiKey) });
            return next.handle(modreq);
        }
    }
}
AuthInterceptorService.ɵfac = function AuthInterceptorService_Factory(t) { return new (t || AuthInterceptorService)(_angular_core__WEBPACK_IMPORTED_MODULE_2__["ɵɵinject"](_auth_service__WEBPACK_IMPORTED_MODULE_1__.AuthService)); };
AuthInterceptorService.ɵprov = /*@__PURE__*/ _angular_core__WEBPACK_IMPORTED_MODULE_2__["ɵɵdefineInjectable"]({ token: AuthInterceptorService, factory: AuthInterceptorService.ɵfac });


/***/ }),

/***/ 384:
/*!**************************************!*\
  !*** ./src/app/auth/auth.service.ts ***!
  \**************************************/
/***/ ((__unused_webpack_module, __webpack_exports__, __webpack_require__) => {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   "AuthService": () => (/* binding */ AuthService)
/* harmony export */ });
/* harmony import */ var _angular_common_http__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! @angular/common/http */ 8784);
/* harmony import */ var rxjs__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! rxjs */ 228);
/* harmony import */ var _environments_environment__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ../../environments/environment */ 2340);
/* harmony import */ var _angular_core__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! @angular/core */ 3184);
/* harmony import */ var _angular_router__WEBPACK_IMPORTED_MODULE_4__ = __webpack_require__(/*! @angular/router */ 2816);






class AuthService {
    constructor(http, router) {
        this.http = http;
        this.router = router;
        this.authData = null;
        this.AuthErrorSub = new rxjs__WEBPACK_IMPORTED_MODULE_1__.Subject();
        this.AuthResultSub = new rxjs__WEBPACK_IMPORTED_MODULE_1__.Subject();
        this.SfErrorSub = new rxjs__WEBPACK_IMPORTED_MODULE_1__.Subject();
        this.SfResultSub = new rxjs__WEBPACK_IMPORTED_MODULE_1__.Subject();
    }
    SignUp(UserEmail, Name, Password) {
        const signup = {
            Email: UserEmail,
            Name: encodeURI(Name),
            Password: encodeURI(Password),
            ReturnSecureToken: true,
        };
        this.authObs = this.http.post(_environments_environment__WEBPACK_IMPORTED_MODULE_0__.environment.SignUpUrl, signup);
        this.RequestSub();
    }
    SignIn(UserEmail, Password) {
        const signin = {
            Email: UserEmail,
            Password: encodeURI(Password),
            ReturnSecureToken: true,
        };
        this.authObs = this.http.post(_environments_environment__WEBPACK_IMPORTED_MODULE_0__.environment.SignInUrl, signin);
        this.RequestSub();
    }
    SecondFactorCheck(Passkey) {
        const httpOptions = {
            headers: new _angular_common_http__WEBPACK_IMPORTED_MODULE_2__.HttpHeaders({
                Passcode: Passkey
            })
        };
        this.SecFactorObs = this.http.post(_environments_environment__WEBPACK_IMPORTED_MODULE_0__.environment.TOTPCheckUrl, null, httpOptions);
        this.SecFactorObs.subscribe(response => {
            if (response.Error.Code === 200) {
                this.authData.SecondFactor.CheckResult = true;
            }
            else {
                this.authData.SecondFactor.CheckResult = false;
            }
            localStorage.setItem('userData', JSON.stringify(this.authData));
            this.SfResultSub.next(this.authData.SecondFactor.CheckResult);
            this.SfErrorSub.next(response);
        }, error => {
            const errresp = error.error;
            this.authData.SecondFactor.CheckResult = false;
            this.SfResultSub.next(false);
            this.SfErrorSub.next(errresp);
        });
    }
    SignOut() {
        this.authData = null;
        if (this.autoRefreshToken) {
            clearTimeout(this.autoRefreshToken);
        }
        localStorage.removeItem('userData');
        this.AuthResultSub.next(false);
        this.SfResultSub.next(false);
        this.router.navigate(['/auth']);
    }
    AutoSignIn() {
        const userData = localStorage.getItem('userData');
        if (!userData) {
            return;
        }
        else {
            this.authData = JSON.parse(userData);
            const ExpDur = new Date(this.authData.ExpirationDate).getTime() -
                new Date().getTime();
            this.AuthResultSub.next(true);
            this.AutoSignOut(ExpDur);
        }
    }
    AutoSignOut(ExpiresIn) {
        this.autoRefreshToken = setTimeout(() => {
            this.SignOut();
        }, ExpiresIn);
    }
    RequestSub() {
        this.authObs.subscribe(response => {
            this.authData = response;
            this.authData.ExpirationDate = String(new Date(new Date().getTime() + +this.authData.ExpiresIn * 1000));
            localStorage.setItem('userData', JSON.stringify(this.authData));
            this.AuthResultSub.next(response.Registered);
            this.AutoSignOut(+this.authData.ExpiresIn * 1000);
        }, error => {
            const errresp = error.error;
            this.AuthResultSub.next(false);
            this.AuthErrorSub.next(errresp);
        });
    }
    CheckRegistered() {
        if (this.authData !== null) {
            if (this.authData.SecondFactor.Enabled === true) {
                return this.authData.SecondFactor.CheckResult && this.authData.Registered;
            }
            else {
                return this.authData.Registered;
            }
        }
        else {
            return false;
        }
    }
    CheckFirstFactorPassed() {
        if (this.authData !== null) {
            return this.authData.Registered;
        }
        else {
            return false;
        }
    }
    HaveToCheckSecondFactor() {
        if (this.authData !== null) {
            if (this.authData.SecondFactor.Enabled) {
                return !this.authData.SecondFactor.CheckResult;
            }
            else {
                return false;
            }
        }
        else {
            return true;
        }
    }
    CheckTokenExpired() {
        if (this.authData.ExpirationDate !== '' && this.authData.ExpirationDate !== null) {
            return !(new Date() > new Date(this.authData.ExpirationDate));
        }
        else {
            return false;
        }
    }
    GetUserToken() {
        if (this.CheckTokenExpired()) {
            return this.authData.Token;
        }
        else {
            return null;
        }
    }
    GetUserEmail() {
        if (this.authData) {
            return this.authData.Email;
        }
        else {
            return null;
        }
    }
    GetUserRole() {
        if (this.authData) {
            return this.authData.Role;
        }
        else {
            return null;
        }
    }
    CheckIfUserIsAdmin() {
        if (this.GetUserRole() === 'admin_role_CRUD') {
            return true;
        }
        return false;
    }
}
AuthService.ɵfac = function AuthService_Factory(t) { return new (t || AuthService)(_angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵinject"](_angular_common_http__WEBPACK_IMPORTED_MODULE_2__.HttpClient), _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵinject"](_angular_router__WEBPACK_IMPORTED_MODULE_4__.Router)); };
AuthService.ɵprov = /*@__PURE__*/ _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵdefineInjectable"]({ token: AuthService, factory: AuthService.ɵfac, providedIn: 'root' });


/***/ }),

/***/ 3482:
/*!********************************************!*\
  !*** ./src/app/header/header.component.ts ***!
  \********************************************/
/***/ ((__unused_webpack_module, __webpack_exports__, __webpack_require__) => {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   "HeaderComponent": () => (/* binding */ HeaderComponent)
/* harmony export */ });
/* harmony import */ var _angular_core__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! @angular/core */ 3184);
/* harmony import */ var _auth_auth_service__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ../auth/auth.service */ 384);
/* harmony import */ var _angular_router__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! @angular/router */ 2816);
/* harmony import */ var _angular_common__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! @angular/common */ 6362);
/* harmony import */ var _ng_bootstrap_ng_bootstrap__WEBPACK_IMPORTED_MODULE_4__ = __webpack_require__(/*! @ng-bootstrap/ng-bootstrap */ 7544);
/* harmony import */ var _angular_forms__WEBPACK_IMPORTED_MODULE_5__ = __webpack_require__(/*! @angular/forms */ 587);






function HeaderComponent_nav_0_a_11_Template(rf, ctx) { if (rf & 1) {
    _angular_core__WEBPACK_IMPORTED_MODULE_1__["ɵɵelementStart"](0, "a", 17);
    _angular_core__WEBPACK_IMPORTED_MODULE_1__["ɵɵtext"](1, " Shopping list ");
    _angular_core__WEBPACK_IMPORTED_MODULE_1__["ɵɵelementEnd"]();
} }
function HeaderComponent_nav_0_a_13_Template(rf, ctx) { if (rf & 1) {
    _angular_core__WEBPACK_IMPORTED_MODULE_1__["ɵɵelementStart"](0, "a", 18);
    _angular_core__WEBPACK_IMPORTED_MODULE_1__["ɵɵtext"](1, " Admin ");
    _angular_core__WEBPACK_IMPORTED_MODULE_1__["ɵɵelementEnd"]();
} }
function HeaderComponent_nav_0_Template(rf, ctx) { if (rf & 1) {
    const _r5 = _angular_core__WEBPACK_IMPORTED_MODULE_1__["ɵɵgetCurrentView"]();
    _angular_core__WEBPACK_IMPORTED_MODULE_1__["ɵɵelementStart"](0, "nav", 1)(1, "div", 2)(2, "div", 3)(3, "a", 4);
    _angular_core__WEBPACK_IMPORTED_MODULE_1__["ɵɵtext"](4, "Recipe book");
    _angular_core__WEBPACK_IMPORTED_MODULE_1__["ɵɵelementEnd"]()();
    _angular_core__WEBPACK_IMPORTED_MODULE_1__["ɵɵelementStart"](5, "div", 5)(6, "ul", 6)(7, "li", 7)(8, "a", 8);
    _angular_core__WEBPACK_IMPORTED_MODULE_1__["ɵɵtext"](9, " Recipes ");
    _angular_core__WEBPACK_IMPORTED_MODULE_1__["ɵɵelementEnd"]()();
    _angular_core__WEBPACK_IMPORTED_MODULE_1__["ɵɵelementStart"](10, "li", 7);
    _angular_core__WEBPACK_IMPORTED_MODULE_1__["ɵɵtemplate"](11, HeaderComponent_nav_0_a_11_Template, 2, 0, "a", 9);
    _angular_core__WEBPACK_IMPORTED_MODULE_1__["ɵɵelementEnd"]();
    _angular_core__WEBPACK_IMPORTED_MODULE_1__["ɵɵelementStart"](12, "li", 7);
    _angular_core__WEBPACK_IMPORTED_MODULE_1__["ɵɵtemplate"](13, HeaderComponent_nav_0_a_13_Template, 2, 0, "a", 10);
    _angular_core__WEBPACK_IMPORTED_MODULE_1__["ɵɵelementEnd"]()();
    _angular_core__WEBPACK_IMPORTED_MODULE_1__["ɵɵelementStart"](14, "form", 11, 12);
    _angular_core__WEBPACK_IMPORTED_MODULE_1__["ɵɵlistener"]("ngSubmit", function HeaderComponent_nav_0_Template_form_ngSubmit_14_listener() { _angular_core__WEBPACK_IMPORTED_MODULE_1__["ɵɵrestoreView"](_r5); const _r3 = _angular_core__WEBPACK_IMPORTED_MODULE_1__["ɵɵreference"](15); const ctx_r4 = _angular_core__WEBPACK_IMPORTED_MODULE_1__["ɵɵnextContext"](); return ctx_r4.SearchRecipes(_r3); });
    _angular_core__WEBPACK_IMPORTED_MODULE_1__["ɵɵelementStart"](16, "button", 13);
    _angular_core__WEBPACK_IMPORTED_MODULE_1__["ɵɵtext"](17);
    _angular_core__WEBPACK_IMPORTED_MODULE_1__["ɵɵelementEnd"]();
    _angular_core__WEBPACK_IMPORTED_MODULE_1__["ɵɵelementStart"](18, "button", 14);
    _angular_core__WEBPACK_IMPORTED_MODULE_1__["ɵɵlistener"]("click", function HeaderComponent_nav_0_Template_button_click_18_listener() { _angular_core__WEBPACK_IMPORTED_MODULE_1__["ɵɵrestoreView"](_r5); const ctx_r6 = _angular_core__WEBPACK_IMPORTED_MODULE_1__["ɵɵnextContext"](); return ctx_r6.OnLogout(); });
    _angular_core__WEBPACK_IMPORTED_MODULE_1__["ɵɵtext"](19, "Sign out");
    _angular_core__WEBPACK_IMPORTED_MODULE_1__["ɵɵelementEnd"]();
    _angular_core__WEBPACK_IMPORTED_MODULE_1__["ɵɵelement"](20, "input", 15);
    _angular_core__WEBPACK_IMPORTED_MODULE_1__["ɵɵelementStart"](21, "button", 16);
    _angular_core__WEBPACK_IMPORTED_MODULE_1__["ɵɵtext"](22, "Search");
    _angular_core__WEBPACK_IMPORTED_MODULE_1__["ɵɵelementEnd"]()()()()();
} if (rf & 2) {
    const _r3 = _angular_core__WEBPACK_IMPORTED_MODULE_1__["ɵɵreference"](15);
    const ctx_r0 = _angular_core__WEBPACK_IMPORTED_MODULE_1__["ɵɵnextContext"]();
    _angular_core__WEBPACK_IMPORTED_MODULE_1__["ɵɵadvance"](11);
    _angular_core__WEBPACK_IMPORTED_MODULE_1__["ɵɵproperty"]("ngIf", ctx_r0.UserAdmin);
    _angular_core__WEBPACK_IMPORTED_MODULE_1__["ɵɵadvance"](2);
    _angular_core__WEBPACK_IMPORTED_MODULE_1__["ɵɵproperty"]("ngIf", ctx_r0.UserAdmin);
    _angular_core__WEBPACK_IMPORTED_MODULE_1__["ɵɵadvance"](4);
    _angular_core__WEBPACK_IMPORTED_MODULE_1__["ɵɵtextInterpolate"](ctx_r0.UserEmail);
    _angular_core__WEBPACK_IMPORTED_MODULE_1__["ɵɵadvance"](4);
    _angular_core__WEBPACK_IMPORTED_MODULE_1__["ɵɵproperty"]("disabled", _r3.invalid);
} }
class HeaderComponent {
    constructor(auth, router) {
        this.auth = auth;
        this.router = router;
        this.LoggedIn = false;
        this.SecondFactor = false;
    }
    ngOnInit() {
        this.LoggedIn = this.auth.CheckRegistered();
        this.UserEmail = this.auth.GetUserEmail();
        this.UserAdmin = this.auth.CheckIfUserIsAdmin();
        if (!this.auth.HaveToCheckSecondFactor()) {
            this.SecondFactor = true;
        }
        this.LoginSub = this.auth.AuthResultSub.subscribe((loggedin) => {
            this.LoggedIn = loggedin;
            if (!this.auth.HaveToCheckSecondFactor()) {
                this.SecondFactor = true;
            }
            this.UserEmail = this.auth.GetUserEmail();
            this.UserAdmin = this.auth.CheckIfUserIsAdmin();
        });
        this.SecondFactorSub = this.auth.SfResultSub.subscribe((result) => {
            this.SecondFactor = result;
        });
    }
    SearchRecipes(form) {
        if (form.valid) {
            const fvalue = form.value;
            this.router.navigate(['recipes'], { queryParams: { search: encodeURI(fvalue.searchreq) } });
        }
    }
    ngOnDestroy() {
        this.LoginSub.unsubscribe();
        this.SecondFactorSub.unsubscribe();
    }
    OnLogout() {
        this.LoggedIn = false;
        this.UserEmail = '';
        this.UserAdmin = false;
        this.auth.SignOut();
    }
}
HeaderComponent.ɵfac = function HeaderComponent_Factory(t) { return new (t || HeaderComponent)(_angular_core__WEBPACK_IMPORTED_MODULE_1__["ɵɵdirectiveInject"](_auth_auth_service__WEBPACK_IMPORTED_MODULE_0__.AuthService), _angular_core__WEBPACK_IMPORTED_MODULE_1__["ɵɵdirectiveInject"](_angular_router__WEBPACK_IMPORTED_MODULE_2__.Router)); };
HeaderComponent.ɵcmp = /*@__PURE__*/ _angular_core__WEBPACK_IMPORTED_MODULE_1__["ɵɵdefineComponent"]({ type: HeaderComponent, selectors: [["app-header"]], decls: 1, vars: 1, consts: [["class", "navbar navbar-expand-lg navbar-light bg-light", "style", "background-color: #e3f2fd;", 4, "ngIf"], [1, "navbar", "navbar-expand-lg", "navbar-light", "bg-light", 2, "background-color", "#e3f2fd"], [1, "container-fluid"], [1, "navbar-header"], ["href", "#", 1, "navbar-brand"], [1, "navbar-expand"], [1, "nav", "nav-tabs"], [1, "nav-item"], ["routerLink", "/recipes", "routerLinkActive", "active", 1, "nav-link"], ["routerLink", "/shopping-list", "routerLinkActive", "active", "class", "nav-link", 4, "ngIf"], ["routerLink", "/admin", "routerLinkActive", "active", "class", "nav-link", 4, "ngIf"], [1, "d-flex", 3, "ngSubmit"], ["searchform", "ngForm"], ["placement", "bottom", "ngbTooltip", "View profile", "type", "button", "routerLinkActive", "active", "routerLink", "/profile", 1, "btn", "btn-outline-secondary"], ["type", "button", 1, "btn", "btn-outline-primary", 3, "click"], ["type", "search", "placeholder", "Search", "ngModel", "", "name", "searchreq", "required", "", 1, "form-control", "me-2"], ["type", "submit", 1, "btn", "btn-outline-primary", 3, "disabled"], ["routerLink", "/shopping-list", "routerLinkActive", "active", 1, "nav-link"], ["routerLink", "/admin", "routerLinkActive", "active", 1, "nav-link"]], template: function HeaderComponent_Template(rf, ctx) { if (rf & 1) {
        _angular_core__WEBPACK_IMPORTED_MODULE_1__["ɵɵtemplate"](0, HeaderComponent_nav_0_Template, 23, 4, "nav", 0);
    } if (rf & 2) {
        _angular_core__WEBPACK_IMPORTED_MODULE_1__["ɵɵproperty"]("ngIf", ctx.LoggedIn && ctx.SecondFactor);
    } }, directives: [_angular_common__WEBPACK_IMPORTED_MODULE_3__.NgIf, _ng_bootstrap_ng_bootstrap__WEBPACK_IMPORTED_MODULE_4__.NgbNavbar, _angular_router__WEBPACK_IMPORTED_MODULE_2__.RouterLinkWithHref, _angular_router__WEBPACK_IMPORTED_MODULE_2__.RouterLinkActive, _angular_forms__WEBPACK_IMPORTED_MODULE_5__["ɵNgNoValidate"], _angular_forms__WEBPACK_IMPORTED_MODULE_5__.NgControlStatusGroup, _angular_forms__WEBPACK_IMPORTED_MODULE_5__.NgForm, _ng_bootstrap_ng_bootstrap__WEBPACK_IMPORTED_MODULE_4__.NgbTooltip, _angular_router__WEBPACK_IMPORTED_MODULE_2__.RouterLink, _angular_forms__WEBPACK_IMPORTED_MODULE_5__.DefaultValueAccessor, _angular_forms__WEBPACK_IMPORTED_MODULE_5__.NgControlStatus, _angular_forms__WEBPACK_IMPORTED_MODULE_5__.NgModel, _angular_forms__WEBPACK_IMPORTED_MODULE_5__.RequiredValidator], styles: ["\n/*# sourceMappingURL=data:application/json;base64,eyJ2ZXJzaW9uIjozLCJzb3VyY2VzIjpbXSwibmFtZXMiOltdLCJtYXBwaW5ncyI6IiIsImZpbGUiOiJoZWFkZXIuY29tcG9uZW50LmNzcyJ9 */"] });


/***/ }),

/***/ 4466:
/*!*****************************************!*\
  !*** ./src/app/shared/shared.module.ts ***!
  \*****************************************/
/***/ ((__unused_webpack_module, __webpack_exports__, __webpack_require__) => {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   "SharedModule": () => (/* binding */ SharedModule)
/* harmony export */ });
/* harmony import */ var _angular_common__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! @angular/common */ 6362);
/* harmony import */ var _angular_core__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! @angular/core */ 3184);


class SharedModule {
}
SharedModule.ɵfac = function SharedModule_Factory(t) { return new (t || SharedModule)(); };
SharedModule.ɵmod = /*@__PURE__*/ _angular_core__WEBPACK_IMPORTED_MODULE_0__["ɵɵdefineNgModule"]({ type: SharedModule });
SharedModule.ɵinj = /*@__PURE__*/ _angular_core__WEBPACK_IMPORTED_MODULE_0__["ɵɵdefineInjector"]({ imports: [[
            _angular_common__WEBPACK_IMPORTED_MODULE_1__.CommonModule
        ], _angular_common__WEBPACK_IMPORTED_MODULE_1__.CommonModule] });
(function () { (typeof ngJitMode === "undefined" || ngJitMode) && _angular_core__WEBPACK_IMPORTED_MODULE_0__["ɵɵsetNgModuleScope"](SharedModule, { imports: [_angular_common__WEBPACK_IMPORTED_MODULE_1__.CommonModule], exports: [_angular_common__WEBPACK_IMPORTED_MODULE_1__.CommonModule] }); })();


/***/ }),

/***/ 2340:
/*!*****************************************!*\
  !*** ./src/environments/environment.ts ***!
  \*****************************************/
/***/ ((__unused_webpack_module, __webpack_exports__, __webpack_require__) => {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   "environment": () => (/* binding */ environment)
/* harmony export */ });
// This file can be replaced during build by using the `fileReplacements` array.
// `ng build --prod` replaces `environment.ts` with `environment.prod.ts`.
// The list of file replacements can be found in `angular.json`.
const environment = {
    production: false,
    RecipePageSize: 5,
    ShoppingListPageSize: 14,
    AdminUserListPageSize: 11,
    SessionsListPageSize: 11,
    MediaListPageSize: 12,
    ApiKey: 'AIzaSyB3Jr8tp5wotjeS-re9iBSgX2b1zbM0Fx4',
    ConfirmEmailUrl: '/api/ConfirmEmail',
    ResendEmailUrl: '/api/ConfirmEmail/Send',
    ResetPasswordUrl: '/api/PasswordReset',
    SendEmailResetPassUrl: '/api/PasswordReset/Send',
    SignUpUrl: '/api/Accounts/SignUp',
    SignInUrl: '/api/Accounts/SignIn',
    GetSetRecipesUrl: '/api/Recipes',
    GetSetFileUrl: '/api/Files',
    SearchRecipesUrl: '/api/Recipes/Search',
    GetSetShoppingListUrl: '/api/ShoppingList',
    GetSetUsersUrl: '/api/Users',
    GetSetCurrentUserUrl: '/api/Users/Current',
    GetSetSessionsUrl: '/api/Sessions',
    GetTOTPQRCodeUrl: '/api/TOTP/Qr.png',
    TOTPSettingsUrl: '/api/TOTP/Settings',
    TOTPCheckUrl: '/api/TOTP/Check',
    GetAuthenticatorUrl: 'https://play.google.com/store/apps/details?id=com.google.android.apps.authenticator2',
};
/*
 * For easier debugging in development mode, you can import the following file
 * to ignore zone related error stack frames such as `zone.run`, `zoneDelegate.invokeTask`.
 *
 * This import should be commented out in production mode because it will have a negative impact
 * on performance if an error is thrown.
 */
// import 'zone.js/plugins/zone-error';  // Included with Angular CLI.


/***/ }),

/***/ 4431:
/*!*********************!*\
  !*** ./src/main.ts ***!
  \*********************/
/***/ ((__unused_webpack_module, __webpack_exports__, __webpack_require__) => {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _angular_platform_browser__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! @angular/platform-browser */ 318);
/* harmony import */ var _angular_core__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! @angular/core */ 3184);
/* harmony import */ var _app_app_module__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ./app/app.module */ 6747);
/* harmony import */ var _environments_environment__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! ./environments/environment */ 2340);




if (_environments_environment__WEBPACK_IMPORTED_MODULE_1__.environment.production) {
    (0,_angular_core__WEBPACK_IMPORTED_MODULE_2__.enableProdMode)();
}
_angular_platform_browser__WEBPACK_IMPORTED_MODULE_3__.platformBrowser().bootstrapModule(_app_app_module__WEBPACK_IMPORTED_MODULE_0__.AppModule)
    .catch(err => console.error(err));


/***/ })

},
/******/ __webpack_require__ => { // webpackRuntimeModules
/******/ var __webpack_exec__ = (moduleId) => (__webpack_require__(__webpack_require__.s = moduleId))
/******/ __webpack_require__.O(0, ["vendor"], () => (__webpack_exec__(4431)));
/******/ var __webpack_exports__ = __webpack_require__.O();
/******/ }
]);
//# sourceMappingURL=main.js.map