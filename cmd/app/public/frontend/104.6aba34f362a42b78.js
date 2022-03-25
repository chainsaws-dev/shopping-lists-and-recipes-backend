"use strict";(self.webpackChunkshopping_lists_and_recipes=self.webpackChunkshopping_lists_and_recipes||[]).push([[104],{9104:(J,u,i)=>{i.r(u),i.d(u,{ConfirmEmailModule:()=>A,httpTranslateLoader:()=>f});var g=i(1048),m=i(1404),p=i(2340),t=i(9724),h=i(3649),l=i(4190),v=i(4800),a=i(1659),d=i(2719),x=i(8177);function b(o,s){1&o&&(t._UZ(0,"input",18),t.ALo(1,"translate")),2&o&&t.Q6J("ngbTooltip",t.lcZ(1,1,"RegisteredEmailRequired"))}function _(o,s){1&o&&(t._UZ(0,"input",19),t.ALo(1,"translate"),t.ALo(2,"translate")),2&o&&t.Q6J("placeholder",t.lcZ(1,2,"NewPassword"))("ngbTooltip",t.lcZ(2,4,"MinSixChars"))}function T(o,s){if(1&o&&(t.TgZ(0,"div",20)(1,"div",21)(2,"button",22),t._uU(3),t.ALo(4,"translate"),t.qZA()()()),2&o){t.oxw();const e=t.MAs(2),n=t.oxw(2);t.xp6(2),t.Q6J("disabled",e.invalid),t.xp6(1),t.Oqu(t.lcZ(4,2,n.ResetPasswordMode&&n.Token?"Save":"Send"))}}function C(o,s){if(1&o){const e=t.EpF();t.TgZ(0,"div",11)(1,"form",12,13),t.NdJ("ngSubmit",function(){t.CHM(e);const r=t.MAs(2);return t.oxw(2).OnSubmitForm(r)}),t.TgZ(3,"div",14),t.YNc(4,b,2,3,"input",15),t.YNc(5,_,3,6,"input",16),t.qZA(),t.YNc(6,T,5,4,"div",17),t.qZA()()}if(2&o){const e=t.oxw(2);t.xp6(4),t.Q6J("ngIf",!e.Token),t.xp6(1),t.Q6J("ngIf",e.ResetPasswordMode&&e.Token),t.xp6(1),t.Q6J("ngIf",!e.Token||e.ResetPasswordMode)}}function w(o,s){if(1&o){const e=t.EpF();t.TgZ(0,"ngb-alert",25),t.NdJ("close",function(){return t.CHM(e),t.oxw(3).ShowMessage=!1}),t._uU(1),t.ALo(2,"FirstLetterUpperPipe"),t.qZA()}if(2&o){const e=t.oxw(3);t.Q6J("type",e.MessageType),t.xp6(1),t.Oqu(t.lcZ(2,2,e.ResponseFromBackend.Error.Message))}}function M(o,s){if(1&o&&(t.TgZ(0,"div",23),t.YNc(1,w,3,4,"ngb-alert",24),t.qZA()),2&o){const e=t.oxw(2);t.xp6(1),t.Q6J("ngIf",e.ShowMessage)}}function S(o,s){if(1&o&&(t.TgZ(0,"div",4)(1,"div",5)(2,"h3"),t._uU(3),t.ALo(4,"translate"),t.qZA()(),t.YNc(5,C,7,3,"div",6),t.TgZ(6,"div",7)(7,"div",8)(8,"a",9),t._uU(9),t.ALo(10,"translate"),t.qZA()()(),t.YNc(11,M,2,1,"div",10),t.qZA()),2&o){const e=t.oxw();t.xp6(3),t.Oqu(t.lcZ(4,4,e.ResetPasswordMode?"ResetPassword":"ConfirmEmail")),t.xp6(2),t.Q6J("ngIf",!e.Token||e.ResetPasswordMode),t.xp6(4),t.Oqu(t.lcZ(10,6,"Back")),t.xp6(2),t.Q6J("ngIf",e.ShowMessage)}}function Z(o,s){1&o&&(t.TgZ(0,"div",26)(1,"span",27),t._uU(2),t.ALo(3,"translate"),t.qZA()()),2&o&&(t.xp6(2),t.Oqu(t.lcZ(3,1,"Loading")))}let E=(()=>{class o{constructor(e,n,r,c,O){this.DataServ=e,this.activeroute=n,this.router=r,this.translate=c,this.sitetitle=O,this.IsLoading=!1,c.addLangs(p.N.SupportedLangs),c.setDefaultLang(p.N.DefaultLocale)}ngOnDestroy(){this.RecivedErrorSub.unsubscribe(),this.RecivedResponseSub.unsubscribe(),this.DataServiceSub.unsubscribe()}ngOnInit(){const e=localStorage.getItem("userLang");this.SwitchLanguage(null!==e?e:p.N.DefaultLocale),this.DataServiceSub=this.DataServ.LoadingData.subscribe(n=>{this.IsLoading=n}),this.RecivedErrorSub=this.DataServ.RecivedError.subscribe(n=>{this.ShowMessage=!0,this.ResponseFromBackend=n,setTimeout(()=>this.ShowMessage=!1,5e3),n&&(this.MessageType=200===n.Error.Code?"success":"danger")}),this.activeroute.queryParams.subscribe(n=>{this.Token=n.Token;const r=this.getUrlWithoutParams();this.ResetPasswordMode="/reset-password"===r,this.ResetPasswordMode||this.Token&&this.DataServ.ConfirmEmail(this.Token)})}SwitchLanguage(e){this.translate.use(e),localStorage.setItem("userLang",e),this.translate.get("WebsiteTitleText",e).subscribe({next:n=>{this.sitetitle.setTitle(n)},error:n=>{console.log(n)}})}getUrlWithoutParams(){const e=this.router.parseUrl(this.router.url);return e.queryParams={},e.toString()}OnSubmitForm(e){e.valid&&(this.IsLoading=!0,this.ResetPasswordMode&&this.Token?this.DataServ.SubmitNewPassword(this.Token,e.value.newpassword):this.ResetPasswordMode?this.DataServ.SendEmailResetPassword(e.value.email):this.DataServ.SendEmailConfirmEmail(e.value.email),e.reset())}}return o.\u0275fac=function(e){return new(e||o)(t.Y36(h.Z),t.Y36(m.gz),t.Y36(m.F0),t.Y36(l.sK),t.Y36(v.Dx))},o.\u0275cmp=t.Xpm({type:o,selectors:[["app-confirm-email"]],decls:4,vars:2,consts:[[1,"main-parent-login"],["class","card mx-auto","style","margin-top: 3px;",4,"ngIf"],[1,"text-center"],["class","spinner-border","role","status",4,"ngIf"],[1,"card","mx-auto",2,"margin-top","3px"],[1,"card-header"],["class","card-body",4,"ngIf"],[1,"card-footer"],[1,"d-flex","justify-content-center"],["href","/"],["style","padding: 3px",4,"ngIf"],[1,"card-body"],[3,"ngSubmit"],["ResendConfEmailForm","ngForm"],[1,"input-group","form-group"],["type","email","name","email","class","form-control mb-3","placeholder","E-mail","placement","right","ngModel","","email","","required","",3,"ngbTooltip",4,"ngIf"],["type","password","name","newpassword","class","form-control mb-3","ngModel","","minlength","6","placement","right","required","",3,"placeholder","ngbTooltip",4,"ngIf"],["class","form-group float-right",4,"ngIf"],["type","email","name","email","placeholder","E-mail","placement","right","ngModel","","email","","required","",1,"form-control","mb-3",3,"ngbTooltip"],["type","password","name","newpassword","ngModel","","minlength","6","placement","right","required","",1,"form-control","mb-3",3,"placeholder","ngbTooltip"],[1,"form-group","float-right"],[1,"input-group"],["type","submit",1,"btn","btn","btn-outline-primary",3,"disabled"],[2,"padding","3px"],[3,"type","close",4,"ngIf"],[3,"type","close"],["role","status",1,"spinner-border"],[1,"visually-hidden-focusable"]],template:function(e,n){1&e&&(t.TgZ(0,"div",0),t.YNc(1,S,12,8,"div",1),t.TgZ(2,"div",2),t.YNc(3,Z,4,3,"div",3),t.qZA()()),2&e&&(t.xp6(1),t.Q6J("ngIf",!n.IsLoading),t.xp6(2),t.Q6J("ngIf",n.IsLoading))},directives:[g.O5,a._Y,a.JL,a.F,a.Fj,a.JJ,a.On,a.on,a.Q7,d._L,a.wO,d.xm],pipes:[l.X$,x.J],styles:[".main-parent-login[_ngcontent-%COMP%]{position:absolute;top:50%;left:50%;transform:translate(-50%,-50%)}.card[_ngcontent-%COMP%]{width:300px}.remember[_ngcontent-%COMP%]   input[_ngcontent-%COMP%]{width:20px;height:20px;margin-left:15px;margin-right:5px}input.ng-invalid.ng-touched[_ngcontent-%COMP%]{border-color:#ff0a1d!important;box-shadow:none!important}input.ng-invalid.ng-touched[_ngcontent-%COMP%]:focus{box-shadow:0 0 0 .2rem #ff0a1d40!important}input.ng-valid.ng-touched[_ngcontent-%COMP%]{border-color:#67c07b!important;box-shadow:none!important}input.ng-valid.ng-touched[_ngcontent-%COMP%]:focus{box-shadow:0 0 0 .2rem #34b30d40!important}"]}),o})();var L=i(9725),y=i(319),P=i(4466);const I=[{path:"",component:E}];let A=(()=>{class o{}return o.\u0275fac=function(e){return new(e||o)},o.\u0275mod=t.oAB({type:o}),o.\u0275inj=t.cJS({imports:[[P.m,g.ez,a.u5,d._A,m.Bz.forChild(I),l.aw.forRoot({loader:{provide:l.Zw,useFactory:f,deps:[L.eN]}}),d.HK]]}),o})();function f(o){return new y.w(o)}}}]);