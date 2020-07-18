(window.webpackJsonp=window.webpackJsonp||[]).push([[9],{CXQP:function(e,t,i){"use strict";i.r(t),i.d(t,"ShoppingListModule",(function(){return P}));var n=i("ofXK"),s=i("AytR"),o=i("fXoL"),r=i("ozzT"),c=i("tyNb"),a=i("GXvH"),h=i("EiSp"),b=i("3Pt+");const l=["f"];function g(e,t){if(1&e){const e=o.Qb();o.Pb(0,"button",14),o.ac("click",(function(){return o.pc(e),o.cc().DeleteSelectedItem()})),o.yc(1,"Delete"),o.Ob()}}function d(e,t){if(1&e){const e=o.Qb();o.Pb(0,"button",15),o.ac("click",(function(){return o.pc(e),o.cc().ClearAllItems()})),o.yc(1,"Clear"),o.Ob()}}let u=(()=>{class e{constructor(e,t){this.ShopListServ=e,this.DataServ=t,this.editmode=!1}ngOnInit(){this.ingselected=this.ShopListServ.IngredientSelected.subscribe(e=>{this.selectedingredient=e,this.editmode=!0,this.slEditForm.setValue({name:this.selectedingredient.Name,amount:this.selectedingredient.Amount})}),this.IngAdd=this.ShopListServ.IngredientAdded.subscribe(e=>{this.DataServ.SaveShoppingList(e)}),this.IngUpd=this.ShopListServ.IngredientUpdated.subscribe(e=>{this.DataServ.SaveShoppingList(e)}),this.IngDel=this.ShopListServ.IngredientDeleted.subscribe(e=>{this.DataServ.DeleteShoppingList(e)}),this.IngCle=this.ShopListServ.IngredientClear.subscribe(()=>{this.DataServ.DeleteAllShoppingList()})}ngOnDestroy(){this.ingselected.unsubscribe(),this.IngAdd.unsubscribe(),this.IngUpd.unsubscribe(),this.IngDel.unsubscribe(),this.IngCle.unsubscribe()}AddNewItem(e){if(e.valid){const t=e.value;this.ShopListServ.AddNewItem(new h.a(t.name,parseInt(t.amount,10)))}}UpdateItem(e){if(e.valid){const t=e.value;this.ShopListServ.UpdateSelectedItem(new h.a(t.name,parseInt(t.amount,10))),this.editmode=!1,this.slEditForm.reset()}}DeleteSelectedItem(){this.ShopListServ.DeleteSelectedItem()}ClearAllItems(){this.ShopListServ.ClearAll()}}return e.\u0275fac=function(t){return new(t||e)(o.Kb(r.a),o.Kb(a.a))},e.\u0275cmp=o.Eb({type:e,selectors:[["app-shopping-edit"]],viewQuery:function(e,t){var i;1&e&&o.Cc(l,!0),2&e&&o.nc(i=o.bc())&&(t.slEditForm=i.first)},decls:18,vars:4,consts:[[1,"row"],[1,"col"],[3,"ngSubmit"],["f","ngForm"],[1,"input-group","mt-3"],[1,"col-sm-9","form-group"],["type","text","id","name","placeholder","Name","name","name","ngModel","","required","",1,"form-control"],[1,"col","form-group"],["type","number","id","amount","placeholder","Amount","name","amount","ngModel","","required","","pattern","^[1-9]+[0-9]*$",1,"form-control"],[1,"input-group"],[1,"input-group-prepend"],["type","submit",1,"btn","btn-outline-primary",3,"disabled"],["class","btn btn-outline-danger","type","button",3,"click",4,"ngIf"],["class","btn btn-outline-secondary","type","button",3,"click",4,"ngIf"],["type","button",1,"btn","btn-outline-danger",3,"click"],["type","button",1,"btn","btn-outline-secondary",3,"click"]],template:function(e,t){if(1&e){const e=o.Qb();o.Pb(0,"div",0),o.Pb(1,"div",1),o.Pb(2,"form",2,3),o.ac("ngSubmit",(function(){o.pc(e);const i=o.oc(3);return t.editmode?t.UpdateItem(i):t.AddNewItem(i)})),o.Pb(4,"div",0),o.Pb(5,"div",4),o.Pb(6,"div",5),o.Lb(7,"input",6),o.Ob(),o.Pb(8,"div",7),o.Lb(9,"input",8),o.Ob(),o.Ob(),o.Ob(),o.Pb(10,"div",0),o.Pb(11,"div",1),o.Pb(12,"div",9),o.Pb(13,"div",10),o.Pb(14,"button",11),o.yc(15),o.Ob(),o.wc(16,g,2,0,"button",12),o.wc(17,d,2,0,"button",13),o.Ob(),o.Ob(),o.Ob(),o.Ob(),o.Ob(),o.Ob(),o.Ob()}if(2&e){const e=o.oc(3);o.xb(14),o.hc("disabled",e.invalid),o.xb(1),o.zc(t.editmode?"Update":"Add"),o.xb(1),o.hc("ngIf",t.ShopListServ.CurrentSelectedItem),o.xb(1),o.hc("ngIf",0!==t.ShopListServ.GetIngredientsLength())}},directives:[b.p,b.h,b.i,b.a,b.g,b.j,b.n,b.l,b.m,n.l],styles:["input.ng-invalid.ng-touched[_ngcontent-%COMP%]{border-color:#ff0a1d!important;box-shadow:none!important}input.ng-invalid.ng-touched[_ngcontent-%COMP%]:focus{box-shadow:0 0 0 .2rem rgba(255,10,29,.25)!important}input.ng-valid.ng-touched[_ngcontent-%COMP%]{border-color:#67c07b!important;box-shadow:none!important}input.ng-valid.ng-touched[_ngcontent-%COMP%]:focus{box-shadow:0 0 0 .2rem rgba(52,179,13,.25)!important}"]}),e})();var p=i("1kSV");function S(e,t){if(1&e){const e=o.Qb();o.Pb(0,"ngb-alert",9),o.ac("close",(function(){return o.pc(e),o.cc(2).ShowMessage=!1})),o.yc(1),o.Ob()}if(2&e){const e=o.cc(2);o.hc("type",e.MessageType),o.xb(1),o.zc(e.ResponseFromBackend.Error.Message)}}function m(e,t){if(1&e){const e=o.Qb();o.Pb(0,"a",10),o.ac("click",(function(){o.pc(e);const i=t.$implicit;return o.cc(2).ShopListServ.SelectItemShopList(i)})),o.yc(1),o.Pb(2,"span",11),o.yc(3),o.Ob(),o.Ob()}if(2&e){const e=t.$implicit,i=o.cc(2);o.hc("ngClass",i.ShopListServ.IsCurrentSelected(e)?"active":""),o.xb(1),o.Ac("",e.Name," "),o.xb(2),o.zc(e.Amount)}}function v(e,t){if(1&e){const e=o.Qb();o.Pb(0,"ngb-pagination",12),o.ac("pageChange",(function(t){return o.pc(e),o.cc(2).OnPageChanged(t)})),o.Ob()}if(2&e){const e=o.cc(2);o.hc("collectionSize",e.slcollectionSize)("pageSize",e.slPageSize)("page",e.slCurrentPage)("rotate",!0)("boundaryLinks",!0)("maxSize",11)}}function f(e,t){if(1&e&&(o.Pb(0,"div",3),o.Pb(1,"div",4),o.Lb(2,"app-shopping-edit"),o.Lb(3,"hr"),o.wc(4,S,2,2,"ngb-alert",5),o.Pb(5,"ul",6),o.wc(6,m,4,3,"a",7),o.Ob(),o.wc(7,v,1,6,"ngb-pagination",8),o.Ob(),o.Ob()),2&e){const e=o.cc();o.xb(4),o.hc("ngIf",e.ShowMessage),o.xb(2),o.hc("ngForOf",e.ingredients),o.xb(1),o.hc("ngIf",e.slcollectionSize>e.slPageSize)}}function I(e,t){1&e&&(o.Pb(0,"div",13),o.Pb(1,"span",14),o.yc(2,"Loading..."),o.Ob(),o.Ob())}const L=[{path:"",redirectTo:"1",pathMatch:"full"},{path:":pn",component:(()=>{class e{constructor(e,t,i,n){this.ShopListServ=e,this.activeroute=t,this.DataServ=i,this.router=n}ngOnDestroy(){this.IngChanged.unsubscribe(),this.PageChanged.unsubscribe(),this.FetchOnInint.unsubscribe(),this.DataLoading.unsubscribe(),this.RecivedErrorSub.unsubscribe(),this.WatchIngAdd.unsubscribe(),this.WatchIngDel.unsubscribe(),this.WatchIngCle.unsubscribe()}ngOnInit(){this.slPageSize=s.a.ShoppingListPageSize,this.RecivedErrorSub=this.DataServ.RecivedError.subscribe(e=>{switch(this.ShowMessage=!0,this.ResponseFromBackend=e,setTimeout(()=>this.ShowMessage=!1,5e3),e.Error.Code){case 200:this.MessageType="success";break;default:this.MessageType="danger"}}),this.IngChanged=this.ShopListServ.IngredientChanged.subscribe(e=>{this.ingredients=e}),this.PageChanged=this.activeroute.params.subscribe(e=>{this.slCurrentPage=+e.pn}),this.DataLoading=this.DataServ.LoadingData.subscribe(e=>{this.IsLoading=e}),this.FetchOnInint=this.DataServ.FetchShoppingList(this.slCurrentPage,s.a.ShoppingListPageSize).subscribe(e=>{this.ingredients=this.ShopListServ.GetIngredients(),this.slcollectionSize=this.ShopListServ.Total},e=>{this.ingredients=[]}),this.WatchIngAdd=this.ShopListServ.IngredientAdded.subscribe(e=>{this.slcollectionSize+=1,this.ingredients=this.ShopListServ.GetIngredients()}),this.WatchIngDel=this.ShopListServ.IngredientDeleted.subscribe(e=>{this.slcollectionSize-=1,this.ingredients=this.ShopListServ.GetIngredients(),0===this.ingredients.length&&(this.slCurrentPage=this.GetPreviousPage(this.slCurrentPage),this.ShopListServ.Total=this.slcollectionSize,0!==this.slcollectionSize&&this.OnPageChanged(this.slCurrentPage))}),this.WatchIngCle=this.ShopListServ.IngredientClear.subscribe(()=>{this.slcollectionSize=0,this.ShopListServ.Total=this.slcollectionSize})}GetPreviousPage(e){return e>1?e-1:1}OnPageChanged(e){this.slCurrentPage=e,this.FetchOnInint=this.DataServ.FetchShoppingList(e,s.a.ShoppingListPageSize).subscribe(()=>{this.ingredients=this.ShopListServ.GetIngredients(),this.router.navigate(["../",e.toString()],{relativeTo:this.activeroute})})}}return e.\u0275fac=function(t){return new(t||e)(o.Kb(r.a),o.Kb(c.a),o.Kb(a.a),o.Kb(c.c))},e.\u0275cmp=o.Eb({type:e,selectors:[["app-shopping-list"]],decls:3,vars:2,consts:[["class","row",4,"ngIf"],[1,"text-center"],["class","spinner-border","role","status",4,"ngIf"],[1,"row"],[1,"col"],[3,"type","close",4,"ngIf"],[1,"list-group"],["style","cursor: pointer;","class","list-group-item list-group-item-action d-flex justify-content-between align-items-center",3,"ngClass","click",4,"ngFor","ngForOf"],[3,"collectionSize","pageSize","page","rotate","boundaryLinks","maxSize","pageChange",4,"ngIf"],[3,"type","close"],[1,"list-group-item","list-group-item-action","d-flex","justify-content-between","align-items-center",2,"cursor","pointer",3,"ngClass","click"],[1,"badge","badge-success","badge-pill"],[3,"collectionSize","pageSize","page","rotate","boundaryLinks","maxSize","pageChange"],["role","status",1,"spinner-border"],[1,"sr-only"]],template:function(e,t){1&e&&(o.wc(0,f,8,3,"div",0),o.Pb(1,"div",1),o.wc(2,I,3,0,"div",2),o.Ob()),2&e&&(o.hc("ngIf",!t.IsLoading),o.xb(2),o.hc("ngIf",t.IsLoading))},directives:[n.l,u,n.k,p.a,n.j,p.i],styles:[""]}),e})()}];let P=(()=>{class e{}return e.\u0275mod=o.Ib({type:e}),e.\u0275inj=o.Hb({factory:function(t){return new(t||e)},imports:[[n.b,b.c,p.b,p.j,c.g.forChild(L)]]}),e})()}}]);