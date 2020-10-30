!function(){function e(e,n){if(!(e instanceof n))throw new TypeError("Cannot call a class as a function")}function n(e,n){for(var t=0;t<n.length;t++){var r=n[t];r.enumerable=r.enumerable||!1,r.configurable=!0,"value"in r&&(r.writable=!0),Object.defineProperty(e,r.key,r)}}function t(e,t,r){return t&&n(e.prototype,t),r&&n(e,r),e}(window.webpackJsonp=window.webpackJsonp||[]).push([[9],{IQtI:function(n,r,o){"use strict";o.r(r),o.d(r,"UserProfileModule",(function(){return k}));var a,c=o("ofXK"),i=o("AytR"),b=o("fXoL"),s=o("tyNb"),u=o("GXvH"),d=o("1kSV"),l=o("3Pt+"),p=o("mrSG"),f=o("tk/3"),h=o("qXBG"),g=((a=function(){function n(t,r){e(this,n),this.http=t,this.auth=r}return t(n,[{key:"transform",value:function(e){return Object(p.a)(this,void 0,void 0,regeneratorRuntime.mark((function n(){var t,r,o,a;return regeneratorRuntime.wrap((function(n){for(;;)switch(n.prev=n.next){case 0:return t=this.auth.GetUserToken(),r=new f.e({Auth:t,ApiKey:i.a.ApiKey}),n.prev=1,n.next=4,this.http.get(e,{headers:r,responseType:"blob"}).toPromise();case 4:return o=n.sent,a=new FileReader,n.abrupt("return",new Promise((function(e,n){a.onloadend=function(){return e(a.result)},a.readAsDataURL(o)})));case 9:return n.prev=9,n.t0=n.catch(1),n.abrupt("return","/favicon.ico");case 12:case"end":return n.stop()}}),n,this,[[1,9]])})))}}]),n}()).\u0275fac=function(e){return new(e||a)(b.Kb(f.b),b.Kb(h.a))},a.\u0275pipe=b.Jb({name:"secureImagePipe",type:a,pure:!0}),a);function m(e,n){1&e&&(b.Pb(0,"div",5),b.Pb(1,"span",6),b.zc(2,"Loading..."),b.Ob(),b.Ob())}function v(e,n){if(1&e){var t=b.Qb();b.Pb(0,"ngb-alert",7),b.ac("close",(function(){return b.qc(t),b.cc().ShowMessage=!1})),b.zc(1),b.Ob()}if(2&e){var r=b.cc();b.hc("type",r.MessageType),b.xb(1),b.Bc(" ",r.ResponseFromBackend.Error.Message,"")}}function w(e,n){if(1&e){var t=b.Qb();b.Pb(0,"form",8,9),b.ac("ngSubmit",(function(){b.qc(t);var e=b.pc(1);return b.cc().OnSaveClick(e)})),b.Pb(2,"div",10),b.Pb(3,"div",11),b.Pb(4,"h2"),b.zc(5,"Edit"),b.Ob(),b.Pb(6,"div",10),b.Pb(7,"label",12),b.zc(8,"Name"),b.Ob(),b.Lb(9,"input",13),b.Ob(),b.Pb(10,"div",10),b.Pb(11,"label",14),b.zc(12,"Email"),b.Ob(),b.Lb(13,"input",15),b.Ob(),b.Pb(14,"div",10),b.Pb(15,"label",16),b.zc(16,"Phone"),b.Ob(),b.Lb(17,"input",17),b.Ob(),b.Pb(18,"label",18),b.zc(19,"Password"),b.Ob(),b.Pb(20,"div",19),b.Pb(21,"div",20),b.Pb(22,"div",21),b.Pb(23,"input",22),b.ac("ngModelChange",(function(e){return b.qc(t),b.cc().changepassword=e})),b.Ob(),b.Ob(),b.Ob(),b.Lb(24,"input",23),b.Ob(),b.Pb(25,"button",24),b.zc(26,"Save"),b.Ob(),b.Pb(27,"button",25),b.zc(28,"Cancel"),b.Ob(),b.Ob(),b.Ob(),b.Ob()}if(2&e){var r=b.pc(1),o=b.cc();b.xb(9),b.hc("ngModel",o.UserToEdit.Name),b.xb(4),b.hc("ngModel",o.UserToEdit.Email),b.xb(4),b.hc("ngModel",o.UserToEdit.Phone),b.xb(6),b.hc("ngModel",o.changepassword),b.xb(1),b.hc("disabled",!o.changepassword),b.xb(1),b.hc("disabled",r.invalid)}}function P(e,n){1&e&&(b.Pb(0,"div",10),b.Pb(1,"h2"),b.zc(2,"Second factor enabled"),b.Ob(),b.Pb(3,"button",26),b.zc(4,"Disable"),b.Ob(),b.Ob())}function y(e,n){if(1&e){var t=b.Qb();b.Pb(0,"form",8,27),b.ac("ngSubmit",(function(){b.qc(t);var e=b.pc(1);return b.cc().OnLinkTwoFactor(e)})),b.Pb(2,"div",10),b.Pb(3,"div",11),b.Pb(4,"h2"),b.zc(5,"Set second factor"),b.Ob(),b.Pb(6,"h4"),b.zc(7,"1. Scan with "),b.Pb(8,"a",28),b.zc(9,"Authenticator"),b.Ob(),b.zc(10,": "),b.Ob(),b.Lb(11,"img",29),b.dc(12,"async"),b.dc(13,"secureImagePipe"),b.Pb(14,"h4"),b.zc(15,"2. Enter the token:"),b.Ob(),b.Pb(16,"label",18),b.zc(17,"Token"),b.Ob(),b.Pb(18,"div",19),b.Pb(19,"div",20),b.Pb(20,"div",21),b.Pb(21,"input",30),b.ac("ngModelChange",(function(e){return b.qc(t),b.cc().TwoFactorEnabled=e})),b.Ob(),b.Ob(),b.Ob(),b.Lb(22,"input",31),b.Ob(),b.Pb(23,"button",32),b.zc(24,"Save"),b.Ob(),b.Ob(),b.Ob(),b.Ob()}if(2&e){var r=b.cc();b.xb(8),b.hc("href",r.AuthUrl,b.rc),b.xb(3),b.hc("src",b.ec(12,4,b.ec(13,6,r.QrUrl)),b.rc),b.xb(10),b.hc("ngModel",r.TwoFactorEnabled),b.xb(1),b.hc("disabled",!r.TwoFactorEnabled)}}var O,x,T=[{path:"",component:(O=function(){function n(t,r,o){e(this,n),this.activatedroute=t,this.router=r,this.datastore=o}return t(n,[{key:"ngOnDestroy",value:function(){}},{key:"ngOnInit",value:function(){var e=this;this.AuthUrl=i.a.GetAuthenticatorUrl,this.QrUrl=i.a.GetTOTPQRCodeUrl,this.FetchUser=this.datastore.CurrentUserFetch.subscribe((function(n){e.UserToEdit=n,e.UserToEdit&&(e.TwoFactorEnabled=e.UserToEdit.SecondFactor)})),this.RecivedErrorSub=this.datastore.RecivedError.subscribe((function(n){switch(e.ShowMessage=!0,e.ResponseFromBackend=n,setTimeout((function(){return e.ShowMessage=!1}),5e3),n.Error.Code){case 200:e.MessageType="success";break;default:e.MessageType="danger"}})),this.DataLoading=this.datastore.LoadingData.subscribe((function(n){e.IsLoading=n})),this.datastore.FetchCurrentUser()}},{key:"OnSaveClick",value:function(e){}},{key:"OnLinkTwoFactor",value:function(e){}}]),n}(),O.\u0275fac=function(e){return new(e||O)(b.Kb(s.a),b.Kb(s.c),b.Kb(u.a))},O.\u0275cmp=b.Eb({type:O,selectors:[["app-user-profile"]],decls:6,vars:5,consts:[[1,"text-center"],["class","spinner-border","role","status",4,"ngIf"],[3,"type","close",4,"ngIf"],[3,"ngSubmit",4,"ngIf"],["class","form-group",4,"ngIf"],["role","status",1,"spinner-border"],[1,"sr-only"],[3,"type","close"],[3,"ngSubmit"],["UserProfileForm","ngForm"],[1,"form-group"],[2,"margin","3px"],["for","name"],["type","text","id","name","placeholder","example","name","username","required","",1,"form-control","mb-1",3,"ngModel"],["for","email"],["type","email","id","email","placeholder","exampe@example.com","name","useremail","required","","email","",1,"form-control","mb-1",3,"ngModel"],["for","phone"],["type","tel","id","phone","placeholder","+7 (965) 777-77-77","name","userphone","pattern","^((8|\\+7)[\\- ]?)?(\\(?\\d{3,4}\\)?[\\- ]?)?[\\d\\- ]{5,10}$",1,"form-control","mb-1",3,"ngModel"],["for","newpassword"],[1,"input-group","mb-3"],[1,"input-group-prepend"],[1,"input-group-text"],["type","checkbox","id","changepassword","name","changepassword",3,"ngModel","ngModelChange"],["type","password","id","newpassword","name","newpassword","placeholder","Enter new password","ngModel","",1,"form-control",3,"disabled"],["type","submit",1,"btn","btn-outline-primary",3,"disabled"],["type","button","routerLink","..",1,"btn","btn-outline-danger",2,"margin-left","3px"],["type","button",1,"btn","btn-outline-danger"],["UserTwoFactorForm","ngForm"],[3,"href"],["alt","QR code",3,"src"],["type","checkbox","id","enabletwofactor","name","enabletwofactor",3,"ngModel","ngModelChange"],["type","password","id","newpasskey","name","passkey","placeholder","Token from Authenticator","ngModel","",1,"form-control",3,"disabled"],["type","button",1,"btn","btn-outline-primary"]],template:function(e,n){1&e&&(b.Pb(0,"div",0),b.xc(1,m,3,0,"div",1),b.Ob(),b.xc(2,v,2,2,"ngb-alert",2),b.xc(3,w,29,6,"form",3),b.xc(4,P,5,0,"div",4),b.xc(5,y,25,8,"form",3)),2&e&&(b.xb(1),b.hc("ngIf",n.IsLoading),b.xb(1),b.hc("ngIf",n.ShowMessage),b.xb(1),b.hc("ngIf",n.UserToEdit),b.xb(1),b.hc("ngIf",n.UserToEdit&&n.TwoFactorEnabled),b.xb(1),b.hc("ngIf",n.UserToEdit&&!n.TwoFactorEnabled))},directives:[c.l,d.a,l.r,l.i,l.j,l.b,l.o,l.h,l.k,l.c,l.n,l.a,s.d],pipes:[c.b,g],styles:[""]}),O)}],k=((x=function n(){e(this,n)}).\u0275mod=b.Ib({type:x}),x.\u0275inj=b.Hb({factory:function(e){return new(e||x)},imports:[[c.c,l.d,d.b,s.g.forChild(T),d.l]]}),x)},mrSG:function(e,n,t){"use strict";function r(e,n,t,r){return new(t||(t=Promise))((function(o,a){function c(e){try{b(r.next(e))}catch(n){a(n)}}function i(e){try{b(r.throw(e))}catch(n){a(n)}}function b(e){var n;e.done?o(e.value):(n=e.value,n instanceof t?n:new t((function(e){e(n)}))).then(c,i)}b((r=r.apply(e,n||[])).next())}))}t.d(n,"a",(function(){return r}))}}])}();