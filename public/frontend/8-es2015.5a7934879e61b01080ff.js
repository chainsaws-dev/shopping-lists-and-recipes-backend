(window.webpackJsonp=window.webpackJsonp||[]).push([[8],{KBmM:function(n,t,e){"use strict";e.r(t),e.d(t,"AuthFeatureModule",(function(){return w}));var i=e("ofXK"),o=e("fXoL"),r=e("qXBG"),s=e("tyNb"),c=e("3Pt+"),a=e("1kSV");function b(n,t){1&n&&(o.Pb(0,"div",9),o.Lb(1,"input",23),o.Ob())}function u(n,t){1&n&&(o.Pb(0,"a",24),o.wc(1,"Forgot your password?"),o.Ob())}function d(n,t){1&n&&(o.Pb(0,"a",25),o.wc(1,"No confirmation email?"),o.Ob())}function g(n,t){if(1&n){const n=o.Qb();o.Pb(0,"ngb-alert",28),o.ac("close",(function(){return o.nc(n),o.cc(3).ShowMessage=!1})),o.wc(1),o.Ob()}if(2&n){const n=o.cc(3);o.fc("type",n.MessageType),o.xb(1),o.xc(n.ResponseFromBackend.Error.Message)}}function p(n,t){if(1&n&&(o.Pb(0,"div",26),o.uc(1,g,2,2,"ngb-alert",27),o.Ob()),2&n){const n=o.cc(2);o.xb(1),o.fc("ngIf",n.ShowMessage)}}function l(n,t){if(1&n){const n=o.Qb();o.Pb(0,"div",4),o.Pb(1,"div",5),o.Pb(2,"h3"),o.wc(3),o.Ob(),o.Ob(),o.Pb(4,"div",6),o.Pb(5,"form",7,8),o.ac("ngSubmit",(function(){o.nc(n);const t=o.mc(6);return o.cc().OnSubmitForm(t)})),o.Pb(7,"div",9),o.Lb(8,"input",10),o.Ob(),o.uc(9,b,2,0,"div",11),o.Pb(10,"div",9),o.Lb(11,"input",12),o.Ob(),o.Pb(12,"div",13),o.Pb(13,"div",14),o.Pb(14,"div",15),o.Pb(15,"button",16),o.wc(16),o.Ob(),o.Pb(17,"button",17),o.ac("click",(function(){return o.nc(n),o.cc().onSwitchMode()})),o.wc(18),o.Ob(),o.Ob(),o.Ob(),o.Ob(),o.Ob(),o.Ob(),o.Pb(19,"div",18),o.Pb(20,"div",19),o.uc(21,u,2,0,"a",20),o.uc(22,d,2,0,"a",21),o.Ob(),o.Ob(),o.uc(23,p,2,1,"div",22),o.Ob()}if(2&n){const n=o.mc(6),t=o.cc();o.xb(3),o.xc(t.LoginMode?"Sign in":"Sign up"),o.xb(6),o.fc("ngIf",!t.LoginMode),o.xb(6),o.fc("disabled",n.invalid),o.xb(1),o.xc(t.LoginMode?"Login":"Sign up"),o.xb(2),o.yc("To ",t.LoginMode?"sign up":"sign in",""),o.xb(3),o.fc("ngIf",t.LoginMode),o.xb(1),o.fc("ngIf",!t.LoginMode),o.xb(1),o.fc("ngIf",t.ShowMessage)}}function h(n,t){1&n&&(o.Pb(0,"div",29),o.Pb(1,"span",30),o.wc(2,"Loading..."),o.Ob(),o.Ob())}let f=(()=>{class n{constructor(n,t){this.authservice=n,this.router=t,this.LoginMode=!0,this.IsLoading=!1}ngOnInit(){this.authservice.CheckRegistered()&&this.Redirect(),this.authErrSub=this.authservice.AuthErrorSub.subscribe(n=>{switch(this.ShowMessage=!0,this.ResponseFromBackend=n,setTimeout(()=>this.ShowMessage=!1,5e3),n.Error.Code){case 200:this.MessageType="success";break;default:this.MessageType="danger"}this.IsLoading=!1}),this.loginResultSub=this.authservice.AuthResultSub.subscribe(n=>{this.loggedIn=n,this.IsLoading=!1,n&&this.Redirect()})}ngOnDestroy(){this.loginResultSub.unsubscribe(),this.authErrSub.unsubscribe()}onSwitchMode(){this.LoginMode=!this.LoginMode}OnSubmitForm(n){this.IsLoading=!0,this.LoginMode?this.authservice.SignIn(n.value.email,n.value.password):this.authservice.SignUp(n.value.email,n.value.name,n.value.password),n.reset()}Redirect(){this.router.navigate(["/recipes"])}}return n.\u0275fac=function(t){return new(t||n)(o.Kb(r.a),o.Kb(s.c))},n.\u0275cmp=o.Eb({type:n,selectors:[["app-auth"]],decls:4,vars:2,consts:[[1,"main-parent-login"],["class","card mx-auto","style","margin-top: 3px;",4,"ngIf"],[1,"text-center"],["class","spinner-border","role","status",4,"ngIf"],[1,"card","mx-auto",2,"margin-top","3px"],[1,"card-header"],[1,"card-body"],[3,"ngSubmit"],["SignupForm","ngForm"],[1,"input-group","form-group"],["type","email","name","email","placeholder","E-mail","ngModel","","email","","required","",1,"form-control"],["class","input-group form-group",4,"ngIf"],["type","password","name","password","placeholder","Password","ngModel","","required","","minlength","6","placement","right","ngbTooltip","Minimum 6 chars",1,"form-control"],[1,"form-group","float-right"],[1,"input-group"],[1,"input-group-prepend"],["type","submit",1,"btn","btn","btn-outline-primary",3,"disabled"],["type","button",1,"btn","btn","btn-outline-primary",3,"click"],[1,"card-footer"],[1,"d-flex","justify-content-center"],["href","#",4,"ngIf"],["href","/confirm-email",4,"ngIf"],["style","padding: 3px",4,"ngIf"],["type","text","name","name","placeholder","Name","ngModel","",1,"form-control"],["href","#"],["href","/confirm-email"],[2,"padding","3px"],[3,"type","close",4,"ngIf"],[3,"type","close"],["role","status",1,"spinner-border"],[1,"sr-only"]],template:function(n,t){1&n&&(o.Pb(0,"div",0),o.uc(1,l,24,8,"div",1),o.Pb(2,"div",2),o.uc(3,h,3,0,"div",3),o.Ob(),o.Ob()),2&n&&(o.xb(1),o.fc("ngIf",!t.IsLoading),o.xb(2),o.fc("ngIf",t.IsLoading))},directives:[i.k,c.r,c.i,c.j,c.b,c.h,c.k,c.c,c.o,c.e,a.a],styles:[".main-parent-login[_ngcontent-%COMP%]{position:absolute;top:50%;left:50%;transform:translate(-50%,-50%)}.card[_ngcontent-%COMP%]{width:300px}.remember[_ngcontent-%COMP%]   input[_ngcontent-%COMP%]{width:20px;height:20px;margin-left:15px;margin-right:5px}input.ng-invalid.ng-touched[_ngcontent-%COMP%]{border-color:#ff0a1d!important;box-shadow:none!important}input.ng-invalid.ng-touched[_ngcontent-%COMP%]:focus{box-shadow:0 0 0 .2rem rgba(255,10,29,.25)!important}input.ng-valid.ng-touched[_ngcontent-%COMP%]{border-color:#67c07b!important;box-shadow:none!important}input.ng-valid.ng-touched[_ngcontent-%COMP%]:focus{box-shadow:0 0 0 .2rem rgba(52,179,13,.25)!important}"]}),n})();var m=e("PCNd");const v=[{path:"",component:f}];let w=(()=>{class n{}return n.\u0275mod=o.Ib({type:n}),n.\u0275inj=o.Hb({factory:function(t){return new(t||n)},imports:[[i.b,c.d,s.g.forChild(v),m.a,a.b]]}),n})()}}]);