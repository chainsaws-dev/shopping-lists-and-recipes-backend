!function(){function e(e,n){if(!(e instanceof n))throw new TypeError("Cannot call a class as a function")}function n(e,n){for(var t=0;t<n.length;t++){var i=n[t];i.enumerable=i.enumerable||!1,i.configurable=!0,"value"in i&&(i.writable=!0),Object.defineProperty(e,i.key,i)}}function t(e,t,i){return t&&n(e.prototype,t),i&&n(e,i),e}(window.webpackJsonp=window.webpackJsonp||[]).push([[9],{jkDv:function(n,i,a){"use strict";a.r(i),a.d(i,"AdminModule",(function(){return ee}));var r,o,s=a("ofXK"),c=a("3Pt+"),u=a("1kSV"),b=a("tyNb"),d=a("P+IX"),l=a("dxYa"),p=a("AytR"),g=a("lJxs"),f=a("fXoL"),m=a("GXvH"),v=((o=function(){function n(t){e(this,n),this.datastorageservice=t}return t(n,[{key:"resolve",value:function(e,n){var t=e.params.id;return this.datastorageservice.FetchUsersList(e.params.pn,p.a.AdminUserListPageSize).pipe(Object(g.a)((function(e){return e.Users[t]})))}}]),n}()).\u0275fac=function(e){return new(e||o)(f.Xb(m.a))},o.\u0275prov=f.Gb({token:o,factory:o.\u0275fac,providedIn:"root"}),o),h=((r=function(){function n(){e(this,n)}return t(n,[{key:"ngOnInit",value:function(){}}]),n}()).\u0275fac=function(e){return new(e||r)},r.\u0275cmp=f.Eb({type:r,selectors:[["app-admin"]],decls:11,vars:0,consts:[[1,"nav","nav-tabs"],[1,"active"],["routerLink","/admin/users","routerLinkActive","active",1,"nav-link"],["routerLink","/admin/sessions","routerLinkActive","active",1,"nav-link"],["routerLink","/admin/media","routerLinkActive","active",1,"nav-link"]],template:function(e,n){1&e&&(f.Pb(0,"ul",0),f.Pb(1,"li",1),f.Pb(2,"a",2),f.xc(3,"Users"),f.Ob(),f.Ob(),f.Pb(4,"li"),f.Pb(5,"a",3),f.xc(6,"Sessions"),f.Ob(),f.Ob(),f.Pb(7,"li"),f.Pb(8,"a",4),f.xc(9,"Media"),f.Ob(),f.Ob(),f.Ob(),f.Lb(10,"router-outlet"))},directives:[b.f,b.e,b.h],styles:[""]}),r),x=a("ggGY"),P=a("c7dm");function O(e,n){1&e&&(f.Pb(0,"div",4),f.Pb(1,"span",5),f.xc(2,"Loading..."),f.Ob(),f.Ob())}function y(e,n){if(1&e){var t=f.Qb();f.Pb(0,"ngb-alert",6),f.ac("close",(function(){return f.oc(t),f.cc().ShowMessage=!1})),f.xc(1),f.Ob()}if(2&e){var i=f.cc();f.fc("type",i.MessageType),f.xb(1),f.zc(" ",i.ResponseFromBackend.Error.Message,"")}}function S(e,n){1&e&&(f.Pb(0,"h2"),f.xc(1,"Edit"),f.Ob())}function w(e,n){1&e&&(f.Pb(0,"h2"),f.xc(1,"Create"),f.Ob())}function k(e,n){if(1&e){var t=f.Qb();f.Pb(0,"form",7,8),f.ac("ngSubmit",(function(){f.oc(t);var e=f.nc(1);return f.cc().OnSaveClick(e)})),f.Pb(2,"div",9),f.Pb(3,"div",10),f.vc(4,S,2,0,"h2",11),f.vc(5,w,2,0,"h2",11),f.Pb(6,"div",9),f.Pb(7,"label",12),f.xc(8,"Name"),f.Ob(),f.Lb(9,"input",13),f.Ob(),f.Pb(10,"div",9),f.Pb(11,"label",14),f.xc(12,"Email"),f.Ob(),f.Lb(13,"input",15),f.Ob(),f.Pb(14,"div",9),f.Pb(15,"label",16),f.xc(16,"Phone"),f.Ob(),f.Lb(17,"input",17),f.Ob(),f.Pb(18,"div",9),f.Pb(19,"label",18),f.xc(20,"Select role:"),f.Ob(),f.Pb(21,"select",19),f.Pb(22,"option",20),f.xc(23,"Guest"),f.Ob(),f.Pb(24,"option",21),f.xc(25,"Admin"),f.Ob(),f.Ob(),f.Ob(),f.Pb(26,"label",22),f.xc(27,"Password"),f.Ob(),f.Pb(28,"div",23),f.Pb(29,"div",24),f.Pb(30,"div",25),f.Pb(31,"input",26),f.ac("ngModelChange",(function(e){return f.oc(t),f.cc().changepassword=e})),f.Ob(),f.Ob(),f.Ob(),f.Lb(32,"input",27),f.Ob(),f.Pb(33,"button",28),f.xc(34,"Save"),f.Ob(),f.Pb(35,"button",29),f.xc(36,"Cancel"),f.Ob(),f.Ob(),f.Ob(),f.Ob()}if(2&e){var i=f.nc(1),a=f.cc();f.xb(4),f.fc("ngIf",a.editmode),f.xb(1),f.fc("ngIf",!a.editmode),f.xb(4),f.fc("ngModel",a.UserToEdit.Name),f.xb(4),f.fc("ngModel",a.UserToEdit.Email),f.xb(4),f.fc("ngModel",a.UserToEdit.Phone),f.xb(4),f.fc("ngModel",a.UserToEdit.Role),f.xb(10),f.fc("ngModel",a.changepassword),f.xb(1),f.fc("disabled",!a.changepassword),f.xb(1),f.fc("disabled",i.invalid)}}var M,C=((M=function(){function n(t,i,a,r){e(this,n),this.AdminServ=t,this.activatedroute=i,this.router=a,this.datastore=r}return t(n,[{key:"ngOnDestroy",value:function(){this.DataLoading.unsubscribe(),this.RecivedErrorSub.unsubscribe(),this.DatabaseUpdated&&this.DatabaseUpdated.unsubscribe()}},{key:"ngOnInit",value:function(){var e=this;this.activatedroute.params.subscribe((function(n){e.editmode=null!=n.id,e.editmode?(e.index=+n.id,e.UserToEdit=e.AdminServ.GetUserById(e.index)):(e.changepassword=!0,e.UserToEdit=new x.a("guest_role_read_only","","","")),e.AdminServ.CurrentSelectedItem=e.UserToEdit})),this.RecivedErrorSub=this.datastore.RecivedError.subscribe((function(n){switch(e.ShowMessage=!0,e.ResponseFromBackend=n,setTimeout((function(){return e.ShowMessage=!1}),5e3),n.Error.Code){case 200:e.MessageType="success";break;default:e.MessageType="danger"}})),this.DataLoading=this.datastore.LoadingData.subscribe((function(n){e.IsLoading=n}))}},{key:"OnSaveClick",value:function(e){var n=this;if(e.valid){if(e.value.changepassword&&0===e.value.newpassword.length)return;this.UserToEdit.Email=e.value.useremail,this.UserToEdit.Name=e.value.username,this.UserToEdit.Phone=e.value.userphone,this.UserToEdit.Role=e.value.roles,this.datastore.SaveUser(this.UserToEdit,e.value.changepassword,e.value.newpassword),this.DatabaseUpdated=this.datastore.UserUpdateInsert.subscribe((function(e){n.editmode?(n.UserToEdit=e,n.AdminServ.UpdateExistingUser(n.UserToEdit,n.index)):(n.UserToEdit=e,n.AdminServ.AddNewUser(n.UserToEdit)),setTimeout((function(){return n.router.navigate(["../"],{relativeTo:n.activatedroute,queryParamsHandling:"merge"})}),1e3)}))}}}]),n}()).\u0275fac=function(e){return new(e||M)(f.Kb(P.a),f.Kb(b.a),f.Kb(b.c),f.Kb(m.a))},M.\u0275cmp=f.Eb({type:M,selectors:[["app-user-edit"]],decls:4,vars:3,consts:[[1,"text-center"],["class","spinner-border","role","status",4,"ngIf"],[3,"type","close",4,"ngIf"],[3,"ngSubmit",4,"ngIf"],["role","status",1,"spinner-border"],[1,"sr-only"],[3,"type","close"],[3,"ngSubmit"],["UserEditForm","ngForm"],[1,"form-group"],[2,"margin","3px"],[4,"ngIf"],["for","name"],["type","text","id","name","placeholder","example","name","username","required","",1,"form-control","mb-1",3,"ngModel"],["for","email"],["type","email","id","email","placeholder","exampe@example.com","name","useremail","required","","email","",1,"form-control","mb-1",3,"ngModel"],["for","phone"],["type","tel","id","phone","placeholder","+7 (965) 777-77-77","name","userphone","pattern","^((8|\\+7)[\\- ]?)?(\\(?\\d{3,4}\\)?[\\- ]?)?[\\d\\- ]{5,10}$",1,"form-control","mb-1",3,"ngModel"],["for","roles"],["name","roles","id","roles","required","",1,"form-control",3,"ngModel"],["value","guest_role_read_only"],["value","admin_role_CRUD"],["for","newpassword"],[1,"input-group","mb-3"],[1,"input-group-prepend"],[1,"input-group-text"],["type","checkbox","id","changepassword","name","changepassword",3,"ngModel","ngModelChange"],["type","password","id","newpassword","name","newpassword","placeholder","Enter new password","ngModel","",1,"form-control",3,"disabled"],["type","submit",1,"btn","btn-outline-primary",3,"disabled"],["type","button","routerLink","..",1,"btn","btn-outline-danger",2,"margin-left","3px"]],template:function(e,n){1&e&&(f.Pb(0,"div",0),f.vc(1,O,3,0,"div",1),f.Ob(),f.vc(2,y,2,2,"ngb-alert",2),f.vc(3,k,37,9,"form",3)),2&e&&(f.xb(1),f.fc("ngIf",n.IsLoading),f.xb(1),f.fc("ngIf",n.ShowMessage),f.xb(1),f.fc("ngIf",n.UserToEdit))},directives:[s.k,u.a,c.r,c.i,c.j,c.b,c.o,c.h,c.k,c.c,c.n,c.p,c.l,c.q,c.a,b.d],styles:["input.ng-invalid.ng-touched[_ngcontent-%COMP%]{border-color:#ff0a1d!important;box-shadow:none!important}input.ng-invalid.ng-touched[_ngcontent-%COMP%]:focus{box-shadow:0 0 0 .2rem rgba(255,10,29,.25)!important}input.ng-valid.ng-touched[_ngcontent-%COMP%]{border-color:#67c07b!important;box-shadow:none!important}input.ng-valid.ng-touched[_ngcontent-%COMP%]:focus{box-shadow:0 0 0 .2rem rgba(52,179,13,.25)!important}select.ng-invalid.ng-touched[_ngcontent-%COMP%]{border-color:#ff0a1d!important;box-shadow:none!important}select.ng-invalid.ng-touched[_ngcontent-%COMP%]:focus{box-shadow:0 0 0 .2rem rgba(255,10,29,.25)!important}select.ng-valid.ng-touched[_ngcontent-%COMP%]{border-color:#67c07b!important;box-shadow:none!important}select.ng-valid.ng-touched[_ngcontent-%COMP%]:focus{box-shadow:0 0 0 .2rem rgba(52,179,13,.25)!important}.image-placeholder[_ngcontent-%COMP%]{background-color:#eee;display:flex;height:333px;margin:5px;border-radius:3px;max-width:592px}.image-placeholder[_ngcontent-%COMP%] > h4[_ngcontent-%COMP%]{align-self:center;text-align:center;width:100%}"]}),M);function I(e,n){if(1&e){var t=f.Qb();f.Pb(0,"ngb-alert",12),f.ac("close",(function(){return f.oc(t),f.cc(2).ShowMessage=!1})),f.xc(1),f.Ob()}if(2&e){var i=f.cc(2);f.fc("type",i.MessageType),f.xb(1),f.zc(" ",i.ResponseFromBackend.Error.Message,"")}}function L(e,n){if(1&e){var t=f.Qb();f.Pb(0,"li",13),f.ac("click",(function(){f.oc(t);var e=n.$implicit;return f.cc(2).AdminServ.SelectItemUsersList(e)})),f.Pb(1,"div",14),f.Pb(2,"p",15),f.xc(3),f.Ob(),f.Pb(4,"div",16),f.Pb(5,"button",17),f.xc(6,"Edit"),f.Ob(),f.Pb(7,"button",18),f.ac("click",(function(){f.oc(t);var e=n.$implicit,i=n.index;return f.cc(2).OnDeleteUser(e,i)})),f.xc(8,"Delete"),f.Ob(),f.Ob(),f.Ob(),f.Ob()}if(2&e){var i=n.$implicit,a=n.index,r=f.cc(2);f.fc("ngClass",r.AdminServ.IsCurrentSelected(i)?"active":""),f.xb(3),f.zc(" ",i.Email,""),f.xb(2),f.fc("routerLink",a)}}function E(e,n){if(1&e){var t=f.Qb();f.Pb(0,"ngb-pagination",19),f.ac("pageChange",(function(e){return f.oc(t),f.cc(2).OnPageChanged(e)})),f.Ob()}if(2&e){var i=f.cc(2);f.fc("collectionSize",i.usCollectionSize)("pageSize",i.usPageSize)("page",i.usCurrentPage)("rotate",!0)("boundaryLinks",!0)("maxSize",11)}}var U=function(e){return["/admin/users",e,"new"]};function D(e,n){if(1&e&&(f.Pb(0,"div",3),f.Pb(1,"div",4),f.Pb(2,"div",5),f.Pb(3,"div",6),f.Pb(4,"button",7),f.xc(5,"Add"),f.Ob(),f.Ob(),f.Ob(),f.vc(6,I,2,2,"ngb-alert",8),f.Pb(7,"ul",9),f.vc(8,L,9,3,"li",10),f.Ob(),f.vc(9,E,1,6,"ngb-pagination",11),f.Ob(),f.Ob()),2&e){var t=f.cc();f.xb(4),f.fc("routerLink",f.jc(4,U,t.usCurrentPage)),f.xb(2),f.fc("ngIf",t.ShowMessage),f.xb(2),f.fc("ngForOf",t.Users),f.xb(1),f.fc("ngIf",t.usCollectionSize>t.usPageSize)}}function z(e,n){1&e&&(f.Pb(0,"div",20),f.Pb(1,"span",21),f.xc(2,"Loading..."),f.Ob(),f.Ob())}var T,A=((T=function(){function n(t,i,a,r){e(this,n),this.ActiveRoute=t,this.DataServ=i,this.AdminServ=a,this.router=r}return t(n,[{key:"ngOnDestroy",value:function(){this.RecivedErrorSub.unsubscribe(),this.PageChanged.unsubscribe(),this.DataLoading.unsubscribe(),this.FetchOnInint.unsubscribe()}},{key:"ngOnInit",value:function(){var e=this;this.usPageSize=p.a.AdminUserListPageSize,this.RecivedErrorSub=this.DataServ.RecivedError.subscribe((function(n){switch(e.ShowMessage=!0,e.ResponseFromBackend=n,setTimeout((function(){return e.ShowMessage=!1}),5e3),n.Error.Code){case 200:e.MessageType="success";break;default:e.MessageType="danger"}})),this.PageChanged=this.ActiveRoute.params.subscribe((function(n){e.usCurrentPage=+n.pn,e.FetchOnInint=e.DataServ.FetchUsersList(e.usCurrentPage,p.a.AdminUserListPageSize).subscribe((function(n){e.Users=e.AdminServ.GetUsers(),e.usCollectionSize=e.AdminServ.Total}),(function(n){e.Users=[]}))})),this.DataLoading=this.DataServ.LoadingData.subscribe((function(n){e.IsLoading=n}))}},{key:"OnPageChanged",value:function(e){this.usCurrentPage=e,this.router.navigate(["../",e.toString()],{relativeTo:this.ActiveRoute})}},{key:"OnDeleteUser",value:function(e,n){this.AdminServ.DeleteUser(n),this.DataServ.DeleteUser(e),this.Users=this.AdminServ.GetUsers()}}]),n}()).\u0275fac=function(e){return new(e||T)(f.Kb(b.a),f.Kb(m.a),f.Kb(P.a),f.Kb(b.c))},T.\u0275cmp=f.Eb({type:T,selectors:[["app-user-list"]],decls:3,vars:2,consts:[["class","row",4,"ngIf"],[1,"text-center"],["class","spinner-border","role","status",4,"ngIf"],[1,"row"],[1,"col"],[1,"input-group","mb-1"],[1,"input-group-prepend","mt-1"],["queryParamsHandling","merge",1,"btn","btn-outline-primary",3,"routerLink"],[3,"type","close",4,"ngIf"],[1,"list-group","mb-1"],["style","cursor: pointer;","class","list-group-item list-group-item-action d-flex justify-content-between align-items-center",3,"ngClass","click",4,"ngFor","ngForOf"],[3,"collectionSize","pageSize","page","rotate","boundaryLinks","maxSize","pageChange",4,"ngIf"],[3,"type","close"],[1,"list-group-item","list-group-item-action","d-flex","justify-content-between","align-items-center",2,"cursor","pointer",3,"ngClass","click"],[1,"input-group"],[1,"p-2","m-0","flex-grow-1",2,"place-items","center"],[1,"input-group-append"],["type","button",1,"btn","btn-outline-success",3,"routerLink"],["type","button",1,"btn","btn-outline-danger",3,"click"],[3,"collectionSize","pageSize","page","rotate","boundaryLinks","maxSize","pageChange"],["role","status",1,"spinner-border"],[1,"sr-only"]],template:function(e,n){1&e&&(f.vc(0,D,10,6,"div",0),f.Pb(1,"div",1),f.vc(2,z,3,0,"div",2),f.Ob()),2&e&&(f.fc("ngIf",!n.IsLoading),f.xb(2),f.fc("ngIf",n.IsLoading))},directives:[s.k,b.d,s.j,u.a,s.i,u.i],styles:[""]}),T),F=a("Ja2n");function R(e,n){if(1&e){var t=f.Qb();f.Pb(0,"ngb-alert",11),f.ac("close",(function(){return f.oc(t),f.cc(2).ShowMessage=!1})),f.xc(1),f.Ob()}if(2&e){var i=f.cc(2);f.fc("type",i.MessageType),f.xb(1),f.zc(" ",i.ResponseFromBackend.Error.Message,"")}}function _(e,n){if(1&e){var t=f.Qb();f.Pb(0,"div",12),f.Pb(1,"a",13),f.Lb(2,"img",14),f.Ob(),f.Pb(3,"div",15),f.Pb(4,"button",16),f.xc(5,"Edit"),f.Ob(),f.Pb(6,"button",17),f.ac("click",(function(){f.oc(t);var e=n.$implicit,i=n.index;return f.cc(2).OnDeleteFile(e,i)})),f.xc(7,"Delete"),f.Ob(),f.Ob(),f.Ob()}if(2&e){var i=n.$implicit,a=n.index;f.xb(1),f.fc("href",i.FileID,f.pc),f.xb(1),f.fc("src",i.FileID,f.pc)("alt",i.Filename),f.xb(2),f.fc("routerLink",a)}}function j(e,n){if(1&e){var t=f.Qb();f.Pb(0,"ngb-pagination",18),f.ac("pageChange",(function(e){return f.oc(t),f.cc(2).OnPageChanged(e)})),f.Ob()}if(2&e){var i=f.cc(2);f.fc("collectionSize",i.meCollectionSize)("pageSize",i.mePageSize)("page",i.meCurrentPage)("rotate",!0)("boundaryLinks",!0)("maxSize",11)}}function K(e,n){if(1&e&&(f.Pb(0,"div"),f.vc(1,R,2,2,"ngb-alert",3),f.Pb(2,"div",4),f.Pb(3,"div",5),f.Pb(4,"button",6),f.xc(5,"Add"),f.Ob(),f.Ob(),f.Ob(),f.Pb(6,"div",7),f.vc(7,_,8,4,"div",8),f.Ob(),f.Pb(8,"div",9),f.vc(9,j,1,6,"ngb-pagination",10),f.Ob(),f.Ob()),2&e){var t=f.cc();f.xb(1),f.fc("ngIf",t.ShowMessage),f.xb(6),f.fc("ngForOf",t.Files),f.xb(2),f.fc("ngIf",t.meCollectionSize>t.mePageSize)}}function G(e,n){1&e&&(f.Pb(0,"div",19),f.Pb(1,"span",20),f.xc(2,"Loading..."),f.Ob(),f.Ob())}var Q,B,q,$,H,N,X=(($=function(){function n(t,i,a,r){e(this,n),this.ActiveRoute=t,this.DataServ=i,this.MediaServ=a,this.router=r}return t(n,[{key:"ngOnDestroy",value:function(){this.RecivedErrorSub.unsubscribe(),this.PageChanged.unsubscribe(),this.DataLoading.unsubscribe(),this.FetchOnInint.unsubscribe()}},{key:"ngOnInit",value:function(){var e=this;this.mePageSize=p.a.MediaListPageSize,this.RecivedErrorSub=this.DataServ.RecivedError.subscribe((function(n){switch(e.ShowMessage=!0,e.ResponseFromBackend=n,setTimeout((function(){return e.ShowMessage=!1}),5e3),n.Error.Code){case 200:e.MessageType="success";break;default:e.MessageType="danger"}})),this.PageChanged=this.ActiveRoute.params.subscribe((function(n){e.meCurrentPage=+n.pn,e.FetchOnInint=e.DataServ.FetchFilesList(e.meCurrentPage,p.a.MediaListPageSize).subscribe((function(n){e.Files=e.MediaServ.GetFiles(),e.meCollectionSize=e.MediaServ.Total}),(function(n){e.Files=[]}))})),this.DataLoading=this.DataServ.LoadingData.subscribe((function(n){e.IsLoading=n}))}},{key:"OnPageChanged",value:function(e){this.meCurrentPage=e,this.router.navigate(["../",e.toString()],{relativeTo:this.ActiveRoute})}},{key:"OnDeleteFile",value:function(e,n){this.MediaServ.DeleteFile(n),this.DataServ.DeleteFile(e.ID,!1),this.Files=this.MediaServ.GetFiles()}}]),n}()).\u0275fac=function(e){return new(e||$)(f.Kb(b.a),f.Kb(m.a),f.Kb(F.a),f.Kb(b.c))},$.\u0275cmp=f.Eb({type:$,selectors:[["app-media-list"]],decls:3,vars:2,consts:[[4,"ngIf"],[1,"text-center"],["class","spinner-border","role","status",4,"ngIf"],[3,"type","close",4,"ngIf"],[1,"input-group","ml-3"],[1,"input-group-prepend","mt-2"],[1,"btn","btn-outline-primary"],[1,"parent-media-box"],["class","media-box",4,"ngFor","ngForOf"],[1,"ml-3"],[3,"collectionSize","pageSize","page","rotate","boundaryLinks","maxSize","pageChange",4,"ngIf"],[3,"type","close"],[1,"media-box"],[3,"href"],[3,"src","alt"],[1,"media-box-overlay"],["type","button",1,"btn","btn-outline-success",3,"routerLink"],["type","button",1,"btn","btn-outline-danger",3,"click"],[3,"collectionSize","pageSize","page","rotate","boundaryLinks","maxSize","pageChange"],["role","status",1,"spinner-border"],[1,"sr-only"]],template:function(e,n){1&e&&(f.vc(0,K,10,3,"div",0),f.Pb(1,"div",1),f.vc(2,G,3,0,"div",2),f.Ob()),2&e&&(f.fc("ngIf",!n.IsLoading),f.xb(2),f.fc("ngIf",n.IsLoading))},directives:[s.k,s.j,u.a,b.d,u.i],styles:[".parent-media-box[_ngcontent-%COMP%]{display:grid;grid-gap:1rem;grid-template-columns:repeat(auto-fit,minmax(360px,1fr));grid-auto-rows:minmax(240px,auto);padding:.3rem 1rem}.media-box[_ngcontent-%COMP%]{font-size:2rem;padding:.3rem;border-radius:5px;display:inline-block;position:relative;border:1px solid rgba(0,0,0,.125);color:#495057;width:100%;height:100%;-o-object-fit:cover;object-fit:cover}.media-box[_ngcontent-%COMP%]   img[_ngcontent-%COMP%]{max-width:100%;max-height:100%;border-radius:5px;display:block}.media-box-overlay[_ngcontent-%COMP%]{position:absolute;bottom:.5rem;right:.5rem}"]}),$),J=((q=function(){function n(){e(this,n)}return t(n,[{key:"ngOnInit",value:function(){}}]),n}()).\u0275fac=function(e){return new(e||q)},q.\u0275cmp=f.Eb({type:q,selectors:[["app-media-edit"]],decls:2,vars:0,template:function(e,n){1&e&&(f.Pb(0,"p"),f.xc(1,"media-edit works!"),f.Ob())},styles:[""]}),q),Y=((B=function(){function n(){e(this,n)}return t(n,[{key:"ngOnInit",value:function(){}}]),n}()).\u0275fac=function(e){return new(e||B)},B.\u0275cmp=f.Eb({type:B,selectors:[["app-sessions-list"]],decls:2,vars:0,template:function(e,n){1&e&&(f.Pb(0,"p"),f.xc(1,"sessions-list works!"),f.Ob())},styles:[""]}),B),V=((Q=function(){function n(){e(this,n)}return t(n,[{key:"ngOnInit",value:function(){}}]),n}()).\u0275fac=function(e){return new(e||Q)},Q.\u0275cmp=f.Eb({type:Q,selectors:[["app-sessions-edit"]],decls:2,vars:0,template:function(e,n){1&e&&(f.Pb(0,"p"),f.xc(1,"sessions-edit works!"),f.Ob())},styles:[""]}),Q),W=[{path:"",redirectTo:"users",pathMatch:"full"},{path:"users",component:h,canActivate:[d.a],children:[{path:"",redirectTo:"1",pathMatch:"full"},{path:":pn",component:A,canActivate:[l.a],data:{expectedRole:"admin_role_CRUD"}},{path:":pn/new",component:C},{path:":pn/:id",component:C,resolve:[v]}]},{path:"sessions",component:h,canActivate:[d.a],children:[{path:"",redirectTo:"1",pathMatch:"full"},{path:":pn",component:Y,canActivate:[l.a],data:{expectedRole:"admin_role_CRUD"}},{path:":pn/new",component:V},{path:":pn/:id",component:V,resolve:[v]}]},{path:"media",component:h,canActivate:[d.a],children:[{path:"",redirectTo:"1",pathMatch:"full"},{path:":pn",component:X,canActivate:[l.a],data:{expectedRole:"admin_role_CRUD"}},{path:":pn/new",component:J},{path:":pn/:id",component:J,resolve:[v]}]}],Z=((N=function n(){e(this,n)}).\u0275mod=f.Ib({type:N}),N.\u0275inj=f.Hb({factory:function(e){return new(e||N)},imports:[[b.g.forChild(W)],b.g]}),N),ee=((H=function n(){e(this,n)}).\u0275mod=f.Ib({type:H}),H.\u0275inj=f.Hb({factory:function(e){return new(e||H)},imports:[[s.b,c.d,u.b,u.j,Z]]}),H)}}])}();