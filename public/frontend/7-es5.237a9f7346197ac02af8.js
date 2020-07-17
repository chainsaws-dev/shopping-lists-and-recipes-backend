!function(){function n(n,t){if(!(n instanceof t))throw new TypeError("Cannot call a class as a function")}function t(n,t){for(var e=0;e<t.length;e++){var i=t[e];i.enumerable=i.enumerable||!1,i.configurable=!0,"value"in i&&(i.writable=!0),Object.defineProperty(n,i.key,i)}}function e(n,e,i){return e&&t(n.prototype,e),i&&t(n,i),n}(window.webpackJsonp=window.webpackJsonp||[]).push([[7],{KBmM:function(t,i,o){"use strict";o.r(i),o.d(i,"AuthFeatureModule",(function(){return O}));var r,a=o("ofXK"),c=o("fXoL"),u=o("qXBG"),s=o("tyNb"),b=o("3Pt+"),d=((r=function(){function t(){n(this,t),this.destroydiv=new c.n}return e(t,[{key:"ngOnInit",value:function(){}},{key:"onClose",value:function(){this.destroydiv.emit()}}]),t}()).\u0275fac=function(n){return new(n||r)},r.\u0275cmp=c.Eb({type:r,selectors:[["app-alert"]],inputs:{message:"message"},outputs:{destroydiv:"destroydiv"},decls:7,vars:1,consts:[[1,"alert","alert-danger",2,"margin-bottom","0"],[1,"row"],[1,"col",2,"margin","auto"],[1,"col-2",2,"text-align","right"],[1,"btn","btn-outline-danger",3,"click"]],template:function(n,t){1&n&&(c.Pb(0,"div",0),c.Pb(1,"div",1),c.Pb(2,"div",2),c.yc(3),c.Ob(),c.Pb(4,"div",3),c.Pb(5,"button",4),c.ac("click",(function(){return t.onClose()})),c.yc(6,"x"),c.Ob(),c.Ob(),c.Ob(),c.Ob()),2&n&&(c.xb(3),c.Ac(" ",t.message," "))},styles:[".backdrop[_ngcontent-%COMP%]{position:fixed;top:0;left:0;width:100vw;height:1000vh;background:rgba(0,0,0,.75);z-index:50}.alert-box[_ngcontent-%COMP%]{position:fixed;border-radius:5px;top:30vh;left:20vw;width:60vw;padding:16px;z-index:100;box-shadow:0 2px 8px rgba(0,0,0,.26)}"]}),r);function l(n,t){if(1&n){var e=c.Qb();c.Pb(0,"div",21),c.Pb(1,"app-alert",22),c.ac("close",(function(){return c.pc(e),c.cc(2).onHandleCloseError()})),c.dc(2,"lowercase"),c.Ob(),c.Ob()}if(2&n){var i=c.cc(2);c.xb(1),c.hc("message",c.ec(2,1,i.authError))}}function p(n,t){if(1&n){var e=c.Qb();c.Pb(0,"div",4),c.Pb(1,"div",5),c.Pb(2,"h3"),c.yc(3),c.Ob(),c.Ob(),c.Pb(4,"div",6),c.Pb(5,"form",7,8),c.ac("ngSubmit",(function(){c.pc(e);var n=c.oc(6);return c.cc().OnSubmitForm(n)})),c.Pb(7,"div",9),c.Lb(8,"input",10),c.Ob(),c.Pb(9,"div",9),c.Lb(10,"input",11),c.Ob(),c.Pb(11,"div",12),c.Pb(12,"div",13),c.Pb(13,"div",14),c.Pb(14,"button",15),c.yc(15),c.Ob(),c.Pb(16,"button",16),c.ac("click",(function(){return c.pc(e),c.cc().onSwitchMode()})),c.yc(17),c.Ob(),c.Ob(),c.Ob(),c.Ob(),c.Ob(),c.Ob(),c.Pb(18,"div",17),c.Pb(19,"div",18),c.Pb(20,"a",19),c.yc(21,"Forgot your password?"),c.Ob(),c.Ob(),c.Ob(),c.wc(22,l,3,3,"div",20),c.Ob()}if(2&n){var i=c.oc(6),o=c.cc();c.xb(3),c.zc(o.LoginMode?"Sign in":"Sign up"),c.xb(11),c.hc("disabled",i.invalid),c.xb(1),c.zc(o.LoginMode?"Login":"Sign up"),c.xb(2),c.Ac("Switch to ",o.LoginMode?"sign up":"sign in",""),c.xb(5),c.hc("ngIf",o.authError)}}function g(n,t){1&n&&(c.Pb(0,"div",23),c.Pb(1,"span",24),c.yc(2,"Loading..."),c.Ob(),c.Ob())}var h,f,v=((h=function(){function t(e,i){n(this,t),this.authservice=e,this.router=i,this.LoginMode=!0,this.IsLoading=!1}return e(t,[{key:"ngOnInit",value:function(){var n=this;this.authservice.CheckRegistered()&&this.Redirect(),this.authErrSub=this.authservice.AuthErrorSub.subscribe((function(t){n.authError=t.replace(/_/g," "),n.IsLoading=!1})),this.loginResultSub=this.authservice.AuthResultSub.subscribe((function(t){n.loggedIn=t,n.IsLoading=!1,t&&n.Redirect()}))}},{key:"ngOnDestroy",value:function(){this.loginResultSub.unsubscribe(),this.authErrSub.unsubscribe()}},{key:"onSwitchMode",value:function(){this.LoginMode=!this.LoginMode}},{key:"OnSubmitForm",value:function(n){this.IsLoading=!0,this.LoginMode?this.authservice.SignIn(n.value.email,n.value.password):this.authservice.SignUp(n.value.email,n.value.password),n.reset()}},{key:"Redirect",value:function(){this.router.navigate(["/recipes"])}},{key:"onHandleCloseError",value:function(){this.authError=null}}]),t}()).\u0275fac=function(n){return new(n||h)(c.Kb(u.a),c.Kb(s.c))},h.\u0275cmp=c.Eb({type:h,selectors:[["app-auth"]],decls:4,vars:2,consts:[[1,"main-parent-login"],["class","card mx-auto","style","margin-top: 3px;",4,"ngIf"],[1,"text-center"],["class","spinner-border","role","status",4,"ngIf"],[1,"card","mx-auto",2,"margin-top","3px"],[1,"card-header"],[1,"card-body"],[3,"ngSubmit"],["SignupForm","ngForm"],[1,"input-group","form-group"],["type","email","name","email","placeholder","E-mail","ngModel","","email","","required","",1,"form-control"],["type","password","name","password","placeholder","Password","ngModel","","required","","minlength","6","placement","right","ngbTooltip","Minimum 6 chars",1,"form-control"],[1,"form-group","float-right"],[1,"input-group"],[1,"input-group-prepend"],["type","submit",1,"btn","btn","btn-outline-primary",3,"disabled"],["type","button",1,"btn","btn","btn-outline-primary",3,"click"],[1,"card-footer"],[1,"d-flex","justify-content-center"],["href","#"],["style","padding: 3px",4,"ngIf"],[2,"padding","3px"],[3,"message","close"],["role","status",1,"spinner-border"],[1,"sr-only"]],template:function(n,t){1&n&&(c.Pb(0,"div",0),c.wc(1,p,23,5,"div",1),c.Pb(2,"div",2),c.wc(3,g,3,0,"div",3),c.Ob(),c.Ob()),2&n&&(c.xb(1),c.hc("ngIf",!t.IsLoading),c.xb(2),c.hc("ngIf",t.IsLoading))},directives:[a.l,b.p,b.h,b.i,b.a,b.g,b.j,b.b,b.n,b.d,d],pipes:[a.i],styles:[".main-parent-login[_ngcontent-%COMP%]{position:absolute;top:50%;left:50%;transform:translate(-50%,-50%)}.card[_ngcontent-%COMP%]{width:400px}.remember[_ngcontent-%COMP%]   input[_ngcontent-%COMP%]{width:20px;height:20px;margin-left:15px;margin-right:5px}input.ng-invalid.ng-touched[_ngcontent-%COMP%]{border-color:#ff0a1d!important;box-shadow:none!important}input.ng-invalid.ng-touched[_ngcontent-%COMP%]:focus{box-shadow:0 0 0 .2rem rgba(255,10,29,.25)!important}input.ng-valid.ng-touched[_ngcontent-%COMP%]{border-color:#67c07b!important;box-shadow:none!important}input.ng-valid.ng-touched[_ngcontent-%COMP%]:focus{box-shadow:0 0 0 .2rem rgba(52,179,13,.25)!important}"]}),h),m=o("PCNd"),y=[{path:"",component:v}],O=((f=function t(){n(this,t)}).\u0275mod=c.Ib({type:f}),f.\u0275inj=c.Hb({factory:function(n){return new(n||f)},imports:[[a.b,b.c,s.g.forChild(y),m.a]]}),f)}}])}();