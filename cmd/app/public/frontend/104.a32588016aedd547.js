"use strict";(self.webpackChunkshopping_lists_and_recipes=self.webpackChunkshopping_lists_and_recipes||[]).push([[104],{9104:(A,u,i)=>{i.r(u),i.d(u,{ConfirmEmailModule:()=>I,httpTranslateLoader:()=>f});var g=i(1048),m=i(1404),p=i(2340),t=i(9724),h=i(3649),l=i(4190),s=i(1659),d=i(2719),v=i(8177);function x(n,a){1&n&&(t._UZ(0,"input",18),t.ALo(1,"translate")),2&n&&t.Q6J("ngbTooltip",t.lcZ(1,1,"RegisteredEmailRequired"))}function _(n,a){1&n&&(t._UZ(0,"input",19),t.ALo(1,"translate"),t.ALo(2,"translate")),2&n&&t.Q6J("placeholder",t.lcZ(1,2,"NewPassword"))("ngbTooltip",t.lcZ(2,4,"MinSixChars"))}function b(n,a){if(1&n&&(t.TgZ(0,"div",20)(1,"div",21)(2,"button",22),t._uU(3),t.ALo(4,"translate"),t.qZA()()()),2&n){t.oxw();const e=t.MAs(2),o=t.oxw(2);t.xp6(2),t.Q6J("disabled",e.invalid),t.xp6(1),t.Oqu(t.lcZ(4,2,o.ResetPasswordMode&&o.Token?"Save":"Send"))}}function C(n,a){if(1&n){const e=t.EpF();t.TgZ(0,"div",11)(1,"form",12,13),t.NdJ("ngSubmit",function(){t.CHM(e);const r=t.MAs(2);return t.oxw(2).OnSubmitForm(r)}),t.TgZ(3,"div",14),t.YNc(4,x,2,3,"input",15),t.YNc(5,_,3,6,"input",16),t.qZA(),t.YNc(6,b,5,4,"div",17),t.qZA()()}if(2&n){const e=t.oxw(2);t.xp6(4),t.Q6J("ngIf",!e.Token),t.xp6(1),t.Q6J("ngIf",e.ResetPasswordMode&&e.Token),t.xp6(1),t.Q6J("ngIf",!e.Token||e.ResetPasswordMode)}}function T(n,a){if(1&n){const e=t.EpF();t.TgZ(0,"ngb-alert",25),t.NdJ("close",function(){return t.CHM(e),t.oxw(3).ShowMessage=!1}),t._uU(1),t.ALo(2,"FirstLetterUpperPipe"),t.qZA()}if(2&n){const e=t.oxw(3);t.Q6J("type",e.MessageType),t.xp6(1),t.Oqu(t.lcZ(2,2,e.ResponseFromBackend.Error.Message))}}function w(n,a){if(1&n&&(t.TgZ(0,"div",23),t.YNc(1,T,3,4,"ngb-alert",24),t.qZA()),2&n){const e=t.oxw(2);t.xp6(1),t.Q6J("ngIf",e.ShowMessage)}}function M(n,a){if(1&n&&(t.TgZ(0,"div",4)(1,"div",5)(2,"h3"),t._uU(3),t.ALo(4,"translate"),t.qZA()(),t.YNc(5,C,7,3,"div",6),t.TgZ(6,"div",7)(7,"div",8)(8,"a",9),t._uU(9),t.ALo(10,"translate"),t.qZA()()(),t.YNc(11,w,2,1,"div",10),t.qZA()),2&n){const e=t.oxw();t.xp6(3),t.Oqu(t.lcZ(4,4,e.ResetPasswordMode?"ResetPassword":"ConfirmEmail")),t.xp6(2),t.Q6J("ngIf",!e.Token||e.ResetPasswordMode),t.xp6(4),t.Oqu(t.lcZ(10,6,"Back")),t.xp6(2),t.Q6J("ngIf",e.ShowMessage)}}function S(n,a){1&n&&(t.TgZ(0,"div",26)(1,"span",27),t._uU(2),t.ALo(3,"translate"),t.qZA()()),2&n&&(t.xp6(2),t.Oqu(t.lcZ(3,1,"Loading")))}let Z=(()=>{class n{constructor(e,o,r,c){this.DataServ=e,this.activeroute=o,this.router=r,this.translate=c,this.IsLoading=!1,c.addLangs(p.N.SupportedLangs),c.setDefaultLang(p.N.DefaultLocale)}ngOnDestroy(){this.RecivedErrorSub.unsubscribe(),this.RecivedResponseSub.unsubscribe(),this.DataServiceSub.unsubscribe()}ngOnInit(){const e=localStorage.getItem("userLang");this.SwitchLanguage(null!==e?e:p.N.DefaultLocale),this.DataServiceSub=this.DataServ.LoadingData.subscribe(o=>{this.IsLoading=o}),this.RecivedErrorSub=this.DataServ.RecivedError.subscribe(o=>{this.ShowMessage=!0,this.ResponseFromBackend=o,setTimeout(()=>this.ShowMessage=!1,5e3),o&&(this.MessageType=200===o.Error.Code?"success":"danger")}),this.activeroute.queryParams.subscribe(o=>{this.Token=o.Token;const r=this.getUrlWithoutParams();this.ResetPasswordMode="/reset-password"===r,this.ResetPasswordMode||this.Token&&this.DataServ.ConfirmEmail(this.Token)})}SwitchLanguage(e){this.translate.use(e),localStorage.setItem("userLang",e)}getUrlWithoutParams(){const e=this.router.parseUrl(this.router.url);return e.queryParams={},e.toString()}OnSubmitForm(e){e.valid&&(this.IsLoading=!0,this.ResetPasswordMode&&this.Token?this.DataServ.SubmitNewPassword(this.Token,e.value.newpassword):this.ResetPasswordMode?this.DataServ.SendEmailResetPassword(e.value.email):this.DataServ.SendEmailConfirmEmail(e.value.email),e.reset())}}return n.\u0275fac=function(e){return new(e||n)(t.Y36(h.Z),t.Y36(m.gz),t.Y36(m.F0),t.Y36(l.sK))},n.\u0275cmp=t.Xpm({type:n,selectors:[["app-confirm-email"]],decls:4,vars:2,consts:[[1,"main-parent-login"],["class","card mx-auto","style","margin-top: 3px;",4,"ngIf"],[1,"text-center"],["class","spinner-border","role","status",4,"ngIf"],[1,"card","mx-auto",2,"margin-top","3px"],[1,"card-header"],["class","card-body",4,"ngIf"],[1,"card-footer"],[1,"d-flex","justify-content-center"],["href","/"],["style","padding: 3px",4,"ngIf"],[1,"card-body"],[3,"ngSubmit"],["ResendConfEmailForm","ngForm"],[1,"input-group","form-group"],["type","email","name","email","class","form-control mb-3","placeholder","E-mail","placement","right","ngModel","","email","","required","",3,"ngbTooltip",4,"ngIf"],["type","password","name","newpassword","class","form-control mb-3","ngModel","","minlength","6","placement","right","required","",3,"placeholder","ngbTooltip",4,"ngIf"],["class","form-group float-right",4,"ngIf"],["type","email","name","email","placeholder","E-mail","placement","right","ngModel","","email","","required","",1,"form-control","mb-3",3,"ngbTooltip"],["type","password","name","newpassword","ngModel","","minlength","6","placement","right","required","",1,"form-control","mb-3",3,"placeholder","ngbTooltip"],[1,"form-group","float-right"],[1,"input-group"],["type","submit",1,"btn","btn","btn-outline-primary",3,"disabled"],[2,"padding","3px"],[3,"type","close",4,"ngIf"],[3,"type","close"],["role","status",1,"spinner-border"],[1,"visually-hidden-focusable"]],template:function(e,o){1&e&&(t.TgZ(0,"div",0),t.YNc(1,M,12,8,"div",1),t.TgZ(2,"div",2),t.YNc(3,S,4,3,"div",3),t.qZA()()),2&e&&(t.xp6(1),t.Q6J("ngIf",!o.IsLoading),t.xp6(2),t.Q6J("ngIf",o.IsLoading))},directives:[g.O5,s._Y,s.JL,s.F,s.Fj,s.JJ,s.On,s.on,s.Q7,d._L,s.wO,d.xm],pipes:[l.X$,v.J],styles:[".main-parent-login[_ngcontent-%COMP%]{position:absolute;top:50%;left:50%;transform:translate(-50%,-50%)}.card[_ngcontent-%COMP%]{width:300px}.remember[_ngcontent-%COMP%]   input[_ngcontent-%COMP%]{width:20px;height:20px;margin-left:15px;margin-right:5px}input.ng-invalid.ng-touched[_ngcontent-%COMP%]{border-color:#ff0a1d!important;box-shadow:none!important}input.ng-invalid.ng-touched[_ngcontent-%COMP%]:focus{box-shadow:0 0 0 .2rem #ff0a1d40!important}input.ng-valid.ng-touched[_ngcontent-%COMP%]{border-color:#67c07b!important;box-shadow:none!important}input.ng-valid.ng-touched[_ngcontent-%COMP%]:focus{box-shadow:0 0 0 .2rem #34b30d40!important}"]}),n})();var E=i(9725),L=i(319),y=i(4466);const P=[{path:"",component:Z}];let I=(()=>{class n{}return n.\u0275fac=function(e){return new(e||n)},n.\u0275mod=t.oAB({type:n}),n.\u0275inj=t.cJS({imports:[[y.m,g.ez,s.u5,d._A,m.Bz.forChild(P),l.aw.forRoot({loader:{provide:l.Zw,useFactory:f,deps:[E.eN]}}),d.HK]]}),n})();function f(n){return new L.w(n)}}}]);