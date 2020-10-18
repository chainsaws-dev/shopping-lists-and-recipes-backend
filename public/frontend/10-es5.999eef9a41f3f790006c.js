!function(){function n(n,e){if(!(n instanceof e))throw new TypeError("Cannot call a class as a function")}function e(n,e){for(var t=0;t<e.length;t++){var o=e[t];o.enumerable=o.enumerable||!1,o.configurable=!0,"value"in o&&(o.writable=!0),Object.defineProperty(n,o.key,o)}}(window.webpackJsonp=window.webpackJsonp||[]).push([[10],{KBmM:function(t,o,i){"use strict";i.r(o),i.d(o,"AuthFeatureModule",(function(){return P}));var r=i("ofXK"),a=i("fXoL"),c=i("qXBG"),s=i("tyNb"),u=i("3Pt+"),b=i("1kSV");function d(n,e){1&n&&(a.Pb(0,"div",9),a.Lb(1,"input",23),a.Ob())}function g(n,e){1&n&&(a.Pb(0,"a",24),a.xc(1,"Forgot your password?"),a.Ob())}function p(n,e){1&n&&(a.Pb(0,"a",25),a.xc(1,"No confirmation email?"),a.Ob())}function l(n,e){if(1&n){var t=a.Qb();a.Pb(0,"ngb-alert",28),a.ac("close",(function(){return a.oc(t),a.cc(3).ShowMessage=!1})),a.xc(1),a.Ob()}if(2&n){var o=a.cc(3);a.fc("type",o.MessageType),a.xb(1),a.yc(o.ResponseFromBackend.Error.Message)}}function f(n,e){if(1&n&&(a.Pb(0,"div",26),a.vc(1,l,2,2,"ngb-alert",27),a.Ob()),2&n){var t=a.cc(2);a.xb(1),a.fc("ngIf",t.ShowMessage)}}function h(n,e){if(1&n){var t=a.Qb();a.Pb(0,"div",4),a.Pb(1,"div",5),a.Pb(2,"h3"),a.xc(3),a.Ob(),a.Ob(),a.Pb(4,"div",6),a.Pb(5,"form",7,8),a.ac("ngSubmit",(function(){a.oc(t);var n=a.nc(6);return a.cc().OnSubmitForm(n)})),a.Pb(7,"div",9),a.Lb(8,"input",10),a.Ob(),a.vc(9,d,2,0,"div",11),a.Pb(10,"div",9),a.Lb(11,"input",12),a.Ob(),a.Pb(12,"div",13),a.Pb(13,"div",14),a.Pb(14,"div",15),a.Pb(15,"button",16),a.xc(16),a.Ob(),a.Pb(17,"button",17),a.ac("click",(function(){return a.oc(t),a.cc().onSwitchMode()})),a.xc(18),a.Ob(),a.Ob(),a.Ob(),a.Ob(),a.Ob(),a.Ob(),a.Pb(19,"div",18),a.Pb(20,"div",19),a.vc(21,g,2,0,"a",20),a.vc(22,p,2,0,"a",21),a.Ob(),a.Ob(),a.vc(23,f,2,1,"div",22),a.Ob()}if(2&n){var o=a.nc(6),i=a.cc();a.xb(3),a.yc(i.LoginMode?"Sign in":"Sign up"),a.xb(6),a.fc("ngIf",!i.LoginMode),a.xb(6),a.fc("disabled",o.invalid),a.xb(1),a.yc(i.LoginMode?"Login":"Sign up"),a.xb(2),a.zc("To ",i.LoginMode?"sign up":"sign in",""),a.xb(3),a.fc("ngIf",i.LoginMode),a.xb(1),a.fc("ngIf",!i.LoginMode),a.xb(1),a.fc("ngIf",i.ShowMessage)}}function m(n,e){1&n&&(a.Pb(0,"div",29),a.Pb(1,"span",30),a.xc(2,"Loading..."),a.Ob(),a.Ob())}var v,x,y=((v=function(){function t(e,o){n(this,t),this.authservice=e,this.router=o,this.LoginMode=!0,this.IsLoading=!1}var o,i,r;return o=t,(i=[{key:"ngOnInit",value:function(){var n=this;this.authservice.CheckRegistered()&&this.Redirect(),this.authErrSub=this.authservice.AuthErrorSub.subscribe((function(e){switch(n.ShowMessage=!0,n.ResponseFromBackend=e,setTimeout((function(){return n.ShowMessage=!1}),5e3),e.Error.Code){case 200:n.MessageType="success";break;default:n.MessageType="danger"}n.IsLoading=!1})),this.loginResultSub=this.authservice.AuthResultSub.subscribe((function(e){n.loggedIn=e,n.IsLoading=!1,e&&n.Redirect()}))}},{key:"ngOnDestroy",value:function(){this.loginResultSub.unsubscribe(),this.authErrSub.unsubscribe()}},{key:"onSwitchMode",value:function(){this.LoginMode=!this.LoginMode}},{key:"OnSubmitForm",value:function(n){this.IsLoading=!0,this.LoginMode?this.authservice.SignIn(n.value.email,n.value.password):this.authservice.SignUp(n.value.email,n.value.name,n.value.password),n.reset()}},{key:"Redirect",value:function(){this.router.navigate(["/recipes"])}}])&&e(o.prototype,i),r&&e(o,r),t}()).\u0275fac=function(n){return new(n||v)(a.Kb(c.a),a.Kb(s.c))},v.\u0275cmp=a.Eb({type:v,selectors:[["app-auth"]],decls:4,vars:2,consts:[[1,"main-parent-login"],["class","card mx-auto","style","margin-top: 3px;",4,"ngIf"],[1,"text-center"],["class","spinner-border","role","status",4,"ngIf"],[1,"card","mx-auto",2,"margin-top","3px"],[1,"card-header"],[1,"card-body"],[3,"ngSubmit"],["SignupForm","ngForm"],[1,"input-group","form-group"],["type","email","name","email","placeholder","E-mail","placement","right","ngbTooltip","Real email required","ngModel","","email","","required","",1,"form-control"],["class","input-group form-group",4,"ngIf"],["type","password","name","password","placeholder","Password","ngModel","","required","","minlength","6","placement","right","ngbTooltip","Minimum 6 chars",1,"form-control"],[1,"form-group","float-right"],[1,"input-group"],[1,"input-group-prepend"],["type","submit",1,"btn","btn","btn-outline-primary",3,"disabled"],["type","button",1,"btn","btn","btn-outline-primary",3,"click"],[1,"card-footer"],[1,"d-flex","justify-content-center"],["href","/reset-password",4,"ngIf"],["href","/confirm-email",4,"ngIf"],["style","padding: 3px",4,"ngIf"],["type","text","name","name","placeholder","Name","ngModel","",1,"form-control"],["href","/reset-password"],["href","/confirm-email"],[2,"padding","3px"],[3,"type","close",4,"ngIf"],[3,"type","close"],["role","status",1,"spinner-border"],[1,"sr-only"]],template:function(n,e){1&n&&(a.Pb(0,"div",0),a.vc(1,h,24,8,"div",1),a.Pb(2,"div",2),a.vc(3,m,3,0,"div",3),a.Ob(),a.Ob()),2&n&&(a.xb(1),a.fc("ngIf",!e.IsLoading),a.xb(2),a.fc("ngIf",e.IsLoading))},directives:[r.k,u.r,u.i,u.j,u.b,b.k,u.h,u.k,u.c,u.o,u.e,b.a],styles:[".main-parent-login[_ngcontent-%COMP%]{position:absolute;top:50%;left:50%;transform:translate(-50%,-50%)}.card[_ngcontent-%COMP%]{width:300px}.remember[_ngcontent-%COMP%]   input[_ngcontent-%COMP%]{width:20px;height:20px;margin-left:15px;margin-right:5px}input.ng-invalid.ng-touched[_ngcontent-%COMP%]{border-color:#ff0a1d!important;box-shadow:none!important}input.ng-invalid.ng-touched[_ngcontent-%COMP%]:focus{box-shadow:0 0 0 .2rem rgba(255,10,29,.25)!important}input.ng-valid.ng-touched[_ngcontent-%COMP%]{border-color:#67c07b!important;box-shadow:none!important}input.ng-valid.ng-touched[_ngcontent-%COMP%]:focus{box-shadow:0 0 0 .2rem rgba(52,179,13,.25)!important}"]}),v),O=i("PCNd"),M=[{path:"",component:y}],P=((x=function e(){n(this,e)}).\u0275mod=a.Ib({type:x}),x.\u0275inj=a.Hb({factory:function(n){return new(n||x)},imports:[[r.b,u.d,s.g.forChild(M),O.a,b.b,b.l]]}),x)}}])}();