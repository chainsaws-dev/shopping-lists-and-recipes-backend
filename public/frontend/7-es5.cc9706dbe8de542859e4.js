function _classCallCheck(n,t){if(!(n instanceof t))throw new TypeError("Cannot call a class as a function")}function _defineProperties(n,t){for(var e=0;e<t.length;e++){var i=t[e];i.enumerable=i.enumerable||!1,i.configurable=!0,"value"in i&&(i.writable=!0),Object.defineProperty(n,i.key,i)}}function _createClass(n,t,e){return t&&_defineProperties(n.prototype,t),e&&_defineProperties(n,e),n}(window.webpackJsonp=window.webpackJsonp||[]).push([[7],{KBmM:function(n,t,e){"use strict";e.r(t),e.d(t,"AuthFeatureModule",(function(){return m}));var i,o=e("ofXK"),r=e("fXoL"),a=e("qXBG"),c=e("tyNb"),s=e("3Pt+"),u=((i=function(){function n(){_classCallCheck(this,n),this.destroydiv=new r.n}return _createClass(n,[{key:"ngOnInit",value:function(){}},{key:"onClose",value:function(){this.destroydiv.emit()}}]),n}()).\u0275fac=function(n){return new(n||i)},i.\u0275cmp=r.Eb({type:i,selectors:[["app-alert"]],inputs:{message:"message"},outputs:{destroydiv:"destroydiv"},decls:7,vars:1,consts:[[1,"alert","alert-danger",2,"margin-bottom","0"],[1,"row"],[1,"col",2,"margin","auto"],[1,"col-2",2,"text-align","right"],[1,"btn","btn-outline-danger",3,"click"]],template:function(n,t){1&n&&(r.Pb(0,"div",0),r.Pb(1,"div",1),r.Pb(2,"div",2),r.zc(3),r.Ob(),r.Pb(4,"div",3),r.Pb(5,"button",4),r.ac("click",(function(){return t.onClose()})),r.zc(6,"x"),r.Ob(),r.Ob(),r.Ob(),r.Ob()),2&n&&(r.xb(3),r.Bc(" ",t.message," "))},styles:[".backdrop[_ngcontent-%COMP%]{position:fixed;top:0;left:0;width:100vw;height:1000vh;background:rgba(0,0,0,.75);z-index:50}.alert-box[_ngcontent-%COMP%]{position:fixed;border-radius:5px;top:30vh;left:20vw;width:60vw;padding:16px;z-index:100;box-shadow:0 2px 8px rgba(0,0,0,.26)}"]}),i);function b(n,t){if(1&n){var e=r.Qb();r.Pb(0,"div",21),r.Pb(1,"app-alert",22),r.ac("close",(function(){return r.qc(e),r.cc(2).onHandleCloseError()})),r.dc(2,"lowercase"),r.Ob(),r.Ob()}if(2&n){var i=r.cc(2);r.xb(1),r.hc("message",r.ec(2,1,i.authError))}}function d(n,t){if(1&n){var e=r.Qb();r.Pb(0,"div",4),r.Pb(1,"div",5),r.Pb(2,"h3"),r.zc(3),r.Ob(),r.Ob(),r.Pb(4,"div",6),r.Pb(5,"form",7,8),r.ac("ngSubmit",(function(){r.qc(e);var n=r.pc(6);return r.cc().OnSubmitForm(n)})),r.Pb(7,"div",9),r.Lb(8,"input",10),r.Ob(),r.Pb(9,"div",9),r.Lb(10,"input",11),r.Ob(),r.Pb(11,"div",12),r.Pb(12,"div",13),r.Pb(13,"div",14),r.Pb(14,"button",15),r.zc(15),r.Ob(),r.Pb(16,"button",16),r.ac("click",(function(){return r.qc(e),r.cc().onSwitchMode()})),r.zc(17),r.Ob(),r.Ob(),r.Ob(),r.Ob(),r.Ob(),r.Ob(),r.Pb(18,"div",17),r.Pb(19,"div",18),r.Pb(20,"a",19),r.zc(21,"Forgot your password?"),r.Ob(),r.Ob(),r.Ob(),r.xc(22,b,3,3,"div",20),r.Ob()}if(2&n){var i=r.pc(6),o=r.cc();r.xb(3),r.Ac(o.LoginMode?"Sign in":"Sign up"),r.xb(11),r.hc("disabled",i.invalid),r.xb(1),r.Ac(o.LoginMode?"Login":"Sign up"),r.xb(2),r.Bc("Switch to ",o.LoginMode?"sign up":"sign in",""),r.xb(5),r.hc("ngIf",o.authError)}}function l(n,t){1&n&&(r.Pb(0,"div",23),r.Pb(1,"span",24),r.zc(2,"Loading..."),r.Ob(),r.Ob())}var p,g,h=((p=function(){function n(t,e){_classCallCheck(this,n),this.authservice=t,this.router=e,this.LoginMode=!0,this.IsLoading=!1}return _createClass(n,[{key:"ngOnInit",value:function(){var n=this;this.authservice.CheckRegistered()&&this.Redirect(),this.authErrSub=this.authservice.AuthErrorSub.subscribe((function(t){n.authError=t.replace(/_/g," "),n.IsLoading=!1})),this.loginResultSub=this.authservice.AuthResultSub.subscribe((function(t){n.loggedIn=t,n.IsLoading=!1,t&&n.Redirect()}))}},{key:"ngOnDestroy",value:function(){this.loginResultSub.unsubscribe(),this.authErrSub.unsubscribe()}},{key:"onSwitchMode",value:function(){this.LoginMode=!this.LoginMode}},{key:"OnSubmitForm",value:function(n){this.IsLoading=!0,this.LoginMode?this.authservice.SignIn(n.value.email,n.value.password):this.authservice.SignUp(n.value.email,n.value.password),n.reset()}},{key:"Redirect",value:function(){this.router.navigate(["/recipes"])}},{key:"onHandleCloseError",value:function(){this.authError=null}}]),n}()).\u0275fac=function(n){return new(n||p)(r.Kb(a.a),r.Kb(c.c))},p.\u0275cmp=r.Eb({type:p,selectors:[["app-auth"]],decls:4,vars:2,consts:[[1,"main-parent-login"],["class","card mx-auto","style","margin-top: 3px;",4,"ngIf"],[1,"text-center"],["class","spinner-border","role","status",4,"ngIf"],[1,"card","mx-auto",2,"margin-top","3px"],[1,"card-header"],[1,"card-body"],[3,"ngSubmit"],["SignupForm","ngForm"],[1,"input-group","form-group"],["type","email","name","email","placeholder","E-mail","ngModel","","email","","required","",1,"form-control"],["type","password","name","password","placeholder","Password","ngModel","","required","","minlength","6","placement","right","ngbTooltip","Minimum 6 chars",1,"form-control"],[1,"form-group","float-right"],[1,"input-group"],[1,"input-group-prepend"],["type","submit",1,"btn","btn","btn-outline-primary",3,"disabled"],["type","button",1,"btn","btn","btn-outline-primary",3,"click"],[1,"card-footer"],[1,"d-flex","justify-content-center"],["href","#"],["style","padding: 3px",4,"ngIf"],[2,"padding","3px"],[3,"message","close"],["role","status",1,"spinner-border"],[1,"sr-only"]],template:function(n,t){1&n&&(r.Pb(0,"div",0),r.xc(1,d,23,5,"div",1),r.Pb(2,"div",2),r.xc(3,l,3,0,"div",3),r.Ob(),r.Ob()),2&n&&(r.xb(1),r.hc("ngIf",!t.IsLoading),r.xb(2),r.hc("ngIf",t.IsLoading))},directives:[o.l,s.p,s.h,s.i,s.a,s.g,s.j,s.b,s.n,s.d,u],pipes:[o.i],styles:[".main-parent-login[_ngcontent-%COMP%]{position:absolute;top:50%;left:50%;transform:translate(-50%,-50%)}.card[_ngcontent-%COMP%]{width:400px}.remember[_ngcontent-%COMP%]   input[_ngcontent-%COMP%]{width:20px;height:20px;margin-left:15px;margin-right:5px}input.ng-invalid.ng-touched[_ngcontent-%COMP%]{border-color:#ff0a1d!important;box-shadow:none!important}input.ng-invalid.ng-touched[_ngcontent-%COMP%]:focus{box-shadow:0 0 0 .2rem rgba(255,10,29,.25)!important}input.ng-valid.ng-touched[_ngcontent-%COMP%]{border-color:#67c07b!important;box-shadow:none!important}input.ng-valid.ng-touched[_ngcontent-%COMP%]:focus{box-shadow:0 0 0 .2rem rgba(52,179,13,.25)!important}"]}),p),f=e("PCNd"),v=[{path:"",component:h}],m=((g=function n(){_classCallCheck(this,n)}).\u0275mod=r.Ib({type:g}),g.\u0275inj=r.Hb({factory:function(n){return new(n||g)},imports:[[o.b,s.c,c.g.forChild(v),f.a]]}),g)}}]);