!function(){function n(n,e){if(!(n instanceof e))throw new TypeError("Cannot call a class as a function")}function e(n,e){for(var t=0;t<e.length;t++){var o=e[t];o.enumerable=o.enumerable||!1,o.configurable=!0,"value"in o&&(o.writable=!0),Object.defineProperty(n,o.key,o)}}(self.webpackChunkshopping_lists_and_recipes=self.webpackChunkshopping_lists_and_recipes||[]).push([[866],{1244:function(t,o,i){"use strict";i.r(o),i.d(o,{AuthFeatureModule:function(){return w}});var r=i(1116),a=i(5366),s=i(8845),u=i(1190),c=i(1462),g=i(2393);function p(n,e){1&n&&(a.TgZ(0,"div",9),a._UZ(1,"input",23),a.qZA())}function d(n,e){1&n&&(a.TgZ(0,"a",24),a._uU(1,"Forgot your password?"),a.qZA())}function l(n,e){1&n&&(a.TgZ(0,"a",25),a._uU(1,"No confirmation email?"),a.qZA())}function f(n,e){if(1&n){var t=a.EpF();a.TgZ(0,"ngb-alert",28),a.NdJ("close",function(){return a.CHM(t),a.oxw(3).ShowMessage=!1}),a._uU(1),a.qZA()}if(2&n){var o=a.oxw(3);a.Q6J("type",o.MessageType),a.xp6(1),a.Oqu(o.ResponseFromBackend.Error.Message)}}function h(n,e){if(1&n&&(a.TgZ(0,"div",26),a.YNc(1,f,2,2,"ngb-alert",27),a.qZA()),2&n){var t=a.oxw(2);a.xp6(1),a.Q6J("ngIf",t.ShowMessage)}}function m(n,e){if(1&n){var t=a.EpF();a.TgZ(0,"div",4),a.TgZ(1,"div",5),a.TgZ(2,"h3"),a._uU(3),a.qZA(),a.qZA(),a.TgZ(4,"div",6),a.TgZ(5,"form",7,8),a.NdJ("ngSubmit",function(){a.CHM(t);var n=a.MAs(6);return a.oxw().OnSubmitForm(n)}),a.TgZ(7,"div",9),a._UZ(8,"input",10),a.qZA(),a.YNc(9,p,2,0,"div",11),a.TgZ(10,"div",9),a._UZ(11,"input",12),a.qZA(),a.TgZ(12,"div",13),a.TgZ(13,"div",14),a.TgZ(14,"div",15),a.TgZ(15,"button",16),a._uU(16),a.qZA(),a.TgZ(17,"button",17),a.NdJ("click",function(){return a.CHM(t),a.oxw().onSwitchMode()}),a._uU(18),a.qZA(),a.qZA(),a.qZA(),a.qZA(),a.qZA(),a.qZA(),a.TgZ(19,"div",18),a.TgZ(20,"div",19),a.YNc(21,d,2,0,"a",20),a.YNc(22,l,2,0,"a",21),a.qZA(),a.qZA(),a.YNc(23,h,2,1,"div",22),a.qZA()}if(2&n){var o=a.MAs(6),i=a.oxw();a.xp6(3),a.Oqu(i.LoginMode?"Sign in":"Sign up"),a.xp6(6),a.Q6J("ngIf",!i.LoginMode),a.xp6(6),a.Q6J("disabled",o.invalid),a.xp6(1),a.Oqu(i.LoginMode?"Login":"Sign up"),a.xp6(2),a.hij("To ",i.LoginMode?"sign up":"sign in",""),a.xp6(3),a.Q6J("ngIf",i.LoginMode),a.xp6(1),a.Q6J("ngIf",!i.LoginMode),a.xp6(1),a.Q6J("ngIf",i.ShowMessage)}}function v(n,e){1&n&&(a.TgZ(0,"div",29),a.TgZ(1,"span",30),a._uU(2,"Loading..."),a.qZA(),a.qZA())}var b,Z=function(){var t=function(){function t(e,o){n(this,t),this.authservice=e,this.router=o,this.LoginMode=!0,this.IsLoading=!1}var o,i,r;return o=t,(i=[{key:"ngOnInit",value:function(){var n=this;this.authservice.CheckRegistered()&&this.Redirect(),this.authErrSub=this.authservice.AuthErrorSub.subscribe(function(e){if(n.ShowMessage=!0,n.ResponseFromBackend=e,setTimeout(function(){return n.ShowMessage=!1},5e3),e)switch(e.Error.Code){case 200:n.MessageType="success";break;default:n.MessageType="danger"}n.IsLoading=!1}),this.loginResultSub=this.authservice.AuthResultSub.subscribe(function(e){n.loggedIn=e,n.IsLoading=!1,e&&(n.authservice.HaveToCheckSecondFactor()?n.GoToSecondFactor():n.Redirect())})}},{key:"ngOnDestroy",value:function(){this.loginResultSub.unsubscribe(),this.authErrSub.unsubscribe()}},{key:"onSwitchMode",value:function(){this.LoginMode=!this.LoginMode}},{key:"OnSubmitForm",value:function(n){this.IsLoading=!0,this.LoginMode?this.authservice.SignIn(n.value.email,n.value.password):this.authservice.SignUp(n.value.email,n.value.name,n.value.password),n.reset()}},{key:"Redirect",value:function(){this.router.navigate(["/recipes"])}},{key:"GoToSecondFactor",value:function(){this.router.navigate(["/totp"])}}])&&e(o.prototype,i),r&&e(o,r),t}();return t.\u0275fac=function(n){return new(n||t)(a.Y36(s.e),a.Y36(u.F0))},t.\u0275cmp=a.Xpm({type:t,selectors:[["app-auth"]],decls:4,vars:2,consts:[[1,"main-parent-login"],["class","card mx-auto","style","margin-top: 3px;",4,"ngIf"],[1,"text-center"],["class","spinner-border","role","status",4,"ngIf"],[1,"card","mx-auto",2,"margin-top","3px"],[1,"card-header"],[1,"card-body"],[3,"ngSubmit"],["SignupForm","ngForm"],[1,"input-group","form-group"],["type","email","name","email","placeholder","E-mail","placement","right","ngbTooltip","Real email required","ngModel","","email","","required","",1,"form-control"],["class","input-group form-group",4,"ngIf"],["type","password","name","password","placeholder","Password","ngModel","","required","","minlength","6","placement","right","ngbTooltip","Minimum 6 chars",1,"form-control"],[1,"form-group","float-right"],[1,"input-group"],[1,"input-group-prepend"],["type","submit",1,"btn","btn","btn-outline-primary",3,"disabled"],["type","button",1,"btn","btn","btn-outline-primary",3,"click"],[1,"card-footer"],[1,"d-flex","justify-content-center"],["href","/reset-password",4,"ngIf"],["href","/confirm-email",4,"ngIf"],["style","padding: 3px",4,"ngIf"],["type","text","name","name","placeholder","Name","ngModel","",1,"form-control"],["href","/reset-password"],["href","/confirm-email"],[2,"padding","3px"],[3,"type","close",4,"ngIf"],[3,"type","close"],["role","status",1,"spinner-border"],[1,"sr-only"]],template:function(n,e){1&n&&(a.TgZ(0,"div",0),a.YNc(1,m,24,8,"div",1),a.TgZ(2,"div",2),a.YNc(3,v,3,0,"div",3),a.qZA(),a.qZA()),2&n&&(a.xp6(1),a.Q6J("ngIf",!e.IsLoading),a.xp6(2),a.Q6J("ngIf",e.IsLoading))},directives:[r.O5,c._Y,c.JL,c.F,c.Fj,g._L,c.JJ,c.On,c.on,c.Q7,c.wO,g.xm],styles:[".main-parent-login[_ngcontent-%COMP%]{position:absolute;top:50%;left:50%;transform:translate(-50%,-50%)}.card[_ngcontent-%COMP%]{width:300px}.remember[_ngcontent-%COMP%]   input[_ngcontent-%COMP%]{width:20px;height:20px;margin-left:15px;margin-right:5px}input.ng-invalid.ng-touched[_ngcontent-%COMP%]{border-color:#ff0a1d!important;box-shadow:none!important}input.ng-invalid.ng-touched[_ngcontent-%COMP%]:focus{box-shadow:0 0 0 .2rem rgba(255,10,29,.25)!important}input.ng-valid.ng-touched[_ngcontent-%COMP%]{border-color:#67c07b!important;box-shadow:none!important}input.ng-valid.ng-touched[_ngcontent-%COMP%]:focus{box-shadow:0 0 0 .2rem rgba(52,179,13,.25)!important}"]}),t}(),M=i(5425),x=[{path:"",component:Z}],w=((b=function e(){n(this,e)}).\u0275fac=function(n){return new(n||b)},b.\u0275mod=a.oAB({type:b}),b.\u0275inj=a.cJS({imports:[[r.ez,c.u5,u.Bz.forChild(x),M.m,g._A,g.HK]]}),b)}}])}();