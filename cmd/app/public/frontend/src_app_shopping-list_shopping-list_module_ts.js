"use strict";
(self["webpackChunkshopping_lists_and_recipes"] = self["webpackChunkshopping_lists_and_recipes"] || []).push([["src_app_shopping-list_shopping-list_module_ts"],{

/***/ 51:
/*!************************************************************************!*\
  !*** ./src/app/shopping-list/shopping-edit/shopping-edit.component.ts ***!
  \************************************************************************/
/***/ ((__unused_webpack_module, __webpack_exports__, __webpack_require__) => {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   "ShoppingEditComponent": () => (/* binding */ ShoppingEditComponent)
/* harmony export */ });
/* harmony import */ var _shared_shared_model__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ../../shared/shared.model */ 3481);
/* harmony import */ var _angular_core__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! @angular/core */ 3184);
/* harmony import */ var _shopping_list_service__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! ../shopping-list.service */ 2457);
/* harmony import */ var src_app_shared_data_storage_service__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! src/app/shared/data-storage.service */ 3649);
/* harmony import */ var _angular_forms__WEBPACK_IMPORTED_MODULE_4__ = __webpack_require__(/*! @angular/forms */ 587);
/* harmony import */ var _angular_common__WEBPACK_IMPORTED_MODULE_5__ = __webpack_require__(/*! @angular/common */ 6362);






const _c0 = ["f"];
function ShoppingEditComponent_button_16_Template(rf, ctx) { if (rf & 1) {
    const _r4 = _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵgetCurrentView"]();
    _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵelementStart"](0, "button", 14);
    _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵlistener"]("click", function ShoppingEditComponent_button_16_Template_button_click_0_listener() { _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵrestoreView"](_r4); const ctx_r3 = _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵnextContext"](); return ctx_r3.DeleteSelectedItem(); });
    _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵtext"](1, "Delete");
    _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵelementEnd"]();
} }
function ShoppingEditComponent_button_17_Template(rf, ctx) { if (rf & 1) {
    const _r6 = _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵgetCurrentView"]();
    _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵelementStart"](0, "button", 15);
    _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵlistener"]("click", function ShoppingEditComponent_button_17_Template_button_click_0_listener() { _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵrestoreView"](_r6); const ctx_r5 = _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵnextContext"](); return ctx_r5.ClearAllItems(); });
    _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵtext"](1, "Clear");
    _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵelementEnd"]();
} }
class ShoppingEditComponent {
    constructor(ShopListServ, DataServ) {
        this.ShopListServ = ShopListServ;
        this.DataServ = DataServ;
        this.editmode = false;
    }
    ngOnInit() {
        this.ingselected = this.ShopListServ.IngredientSelected.subscribe((ing) => {
            this.selectedingredient = ing;
            this.editmode = true;
            this.slEditForm.setValue({
                name: this.selectedingredient.Name,
                amount: this.selectedingredient.Amount
            });
        });
        this.IngAdd = this.ShopListServ.IngredientAdded.subscribe((addedIng) => {
            this.DataServ.SaveShoppingList(addedIng);
        });
        this.IngUpd = this.ShopListServ.IngredientUpdated.subscribe((updIng) => {
            this.DataServ.SaveShoppingList(updIng);
        });
        this.IngDel = this.ShopListServ.IngredientDeleted.subscribe((delIng) => {
            this.DataServ.DeleteShoppingList(delIng);
        });
        this.IngCle = this.ShopListServ.IngredientClear.subscribe(() => {
            this.DataServ.DeleteAllShoppingList();
        });
    }
    ngOnDestroy() {
        this.ingselected.unsubscribe();
        this.IngAdd.unsubscribe();
        this.IngUpd.unsubscribe();
        this.IngDel.unsubscribe();
        this.IngCle.unsubscribe();
    }
    AddNewItem(form) {
        if (form.valid) {
            const fvalue = form.value;
            this.ShopListServ.AddNewItem(new _shared_shared_model__WEBPACK_IMPORTED_MODULE_0__.Ingredient(fvalue.name, parseInt(fvalue.amount, 10)), false);
        }
    }
    UpdateItem(form) {
        if (form.valid) {
            const fvalue = form.value;
            this.ShopListServ.UpdateSelectedItem(new _shared_shared_model__WEBPACK_IMPORTED_MODULE_0__.Ingredient(fvalue.name, parseInt(fvalue.amount, 10)));
            this.editmode = false;
            this.slEditForm.reset();
        }
    }
    DeleteSelectedItem() {
        this.ShopListServ.DeleteSelectedItem();
    }
    ClearAllItems() {
        this.ShopListServ.ClearAll();
    }
}
ShoppingEditComponent.ɵfac = function ShoppingEditComponent_Factory(t) { return new (t || ShoppingEditComponent)(_angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵdirectiveInject"](_shopping_list_service__WEBPACK_IMPORTED_MODULE_1__.ShoppingListService), _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵdirectiveInject"](src_app_shared_data_storage_service__WEBPACK_IMPORTED_MODULE_2__.DataStorageService)); };
ShoppingEditComponent.ɵcmp = /*@__PURE__*/ _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵdefineComponent"]({ type: ShoppingEditComponent, selectors: [["app-shopping-edit"]], viewQuery: function ShoppingEditComponent_Query(rf, ctx) { if (rf & 1) {
        _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵviewQuery"](_c0, 5);
    } if (rf & 2) {
        let _t;
        _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵqueryRefresh"](_t = _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵloadQuery"]()) && (ctx.slEditForm = _t.first);
    } }, decls: 18, vars: 4, consts: [[1, "row"], [1, "col"], [3, "ngSubmit"], ["f", "ngForm"], [1, "input-group", "mt-2", "mb-2"], [1, "col-sm-9", "form-group", "me-1"], ["type", "text", "id", "name", "placeholder", "Name", "name", "name", "ngModel", "", "required", "", 1, "form-control"], [1, "col", "form-group"], ["type", "number", "id", "amount", "placeholder", "Amount", "name", "amount", "ngModel", "", "required", "", "pattern", "^[1-9]+[0-9]*$", 1, "form-control"], [1, "input-group"], [1, "input-group-prepend"], ["type", "submit", 1, "btn", "btn-outline-primary", "me-1", 3, "disabled"], ["class", "btn btn-outline-danger me-1", "type", "button", 3, "click", 4, "ngIf"], ["class", "btn btn-outline-secondary", "type", "button", 3, "click", 4, "ngIf"], ["type", "button", 1, "btn", "btn-outline-danger", "me-1", 3, "click"], ["type", "button", 1, "btn", "btn-outline-secondary", 3, "click"]], template: function ShoppingEditComponent_Template(rf, ctx) { if (rf & 1) {
        const _r7 = _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵgetCurrentView"]();
        _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵelementStart"](0, "div", 0)(1, "div", 1)(2, "form", 2, 3);
        _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵlistener"]("ngSubmit", function ShoppingEditComponent_Template_form_ngSubmit_2_listener() { _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵrestoreView"](_r7); const _r0 = _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵreference"](3); return ctx.editmode ? ctx.UpdateItem(_r0) : ctx.AddNewItem(_r0); });
        _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵelementStart"](4, "div", 0)(5, "div", 4)(6, "div", 5);
        _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵelement"](7, "input", 6);
        _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵelementEnd"]();
        _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵelementStart"](8, "div", 7);
        _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵelement"](9, "input", 8);
        _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵelementEnd"]()()();
        _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵelementStart"](10, "div", 0)(11, "div", 1)(12, "div", 9)(13, "div", 10)(14, "button", 11);
        _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵtext"](15);
        _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵelementEnd"]();
        _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵtemplate"](16, ShoppingEditComponent_button_16_Template, 2, 0, "button", 12);
        _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵtemplate"](17, ShoppingEditComponent_button_17_Template, 2, 0, "button", 13);
        _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵelementEnd"]()()()()()()();
    } if (rf & 2) {
        const _r0 = _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵreference"](3);
        _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵadvance"](14);
        _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵproperty"]("disabled", _r0.invalid);
        _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵadvance"](1);
        _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵtextInterpolate"](ctx.editmode ? "Update" : "Add");
        _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵadvance"](1);
        _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵproperty"]("ngIf", ctx.ShopListServ.CurrentSelectedItem);
        _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵadvance"](1);
        _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵproperty"]("ngIf", ctx.ShopListServ.GetIngredientsLength() !== 0);
    } }, directives: [_angular_forms__WEBPACK_IMPORTED_MODULE_4__["ɵNgNoValidate"], _angular_forms__WEBPACK_IMPORTED_MODULE_4__.NgControlStatusGroup, _angular_forms__WEBPACK_IMPORTED_MODULE_4__.NgForm, _angular_forms__WEBPACK_IMPORTED_MODULE_4__.DefaultValueAccessor, _angular_forms__WEBPACK_IMPORTED_MODULE_4__.NgControlStatus, _angular_forms__WEBPACK_IMPORTED_MODULE_4__.NgModel, _angular_forms__WEBPACK_IMPORTED_MODULE_4__.RequiredValidator, _angular_forms__WEBPACK_IMPORTED_MODULE_4__.NumberValueAccessor, _angular_forms__WEBPACK_IMPORTED_MODULE_4__.PatternValidator, _angular_common__WEBPACK_IMPORTED_MODULE_5__.NgIf], styles: ["input.ng-invalid.ng-touched[_ngcontent-%COMP%] {\n  border-color: rgba(255, 10, 29, 1) !important;\n  box-shadow: none !important;\n}\n\ninput.ng-invalid.ng-touched[_ngcontent-%COMP%]:focus {\n  box-shadow: 0 0 0 0.2rem rgba(255, 10, 29, 0.25) !important;\n}\n\ninput.ng-valid.ng-touched[_ngcontent-%COMP%] {\n  border-color: rgb(103, 192, 123) !important;\n  box-shadow: none !important;\n}\n\ninput.ng-valid.ng-touched[_ngcontent-%COMP%]:focus {\n  box-shadow: 0 0 0 0.2rem rgba(52, 179, 13, 0.25) !important;\n}\n/*# sourceMappingURL=data:application/json;base64,eyJ2ZXJzaW9uIjozLCJzb3VyY2VzIjpbInNob3BwaW5nLWVkaXQuY29tcG9uZW50LmNzcyJdLCJuYW1lcyI6W10sIm1hcHBpbmdzIjoiQUFBQTtFQUNFLDZDQUE2QztFQUM3QywyQkFBMkI7QUFDN0I7O0FBRUE7RUFDRSwyREFBMkQ7QUFDN0Q7O0FBRUE7RUFDRSwyQ0FBMkM7RUFDM0MsMkJBQTJCO0FBQzdCOztBQUVBO0VBQ0UsMkRBQTJEO0FBQzdEIiwiZmlsZSI6InNob3BwaW5nLWVkaXQuY29tcG9uZW50LmNzcyIsInNvdXJjZXNDb250ZW50IjpbImlucHV0Lm5nLWludmFsaWQubmctdG91Y2hlZCB7XG4gIGJvcmRlci1jb2xvcjogcmdiYSgyNTUsIDEwLCAyOSwgMSkgIWltcG9ydGFudDtcbiAgYm94LXNoYWRvdzogbm9uZSAhaW1wb3J0YW50O1xufVxuXG5pbnB1dC5uZy1pbnZhbGlkLm5nLXRvdWNoZWQ6Zm9jdXMge1xuICBib3gtc2hhZG93OiAwIDAgMCAwLjJyZW0gcmdiYSgyNTUsIDEwLCAyOSwgMC4yNSkgIWltcG9ydGFudDtcbn1cblxuaW5wdXQubmctdmFsaWQubmctdG91Y2hlZCB7XG4gIGJvcmRlci1jb2xvcjogcmdiKDEwMywgMTkyLCAxMjMpICFpbXBvcnRhbnQ7XG4gIGJveC1zaGFkb3c6IG5vbmUgIWltcG9ydGFudDtcbn1cblxuaW5wdXQubmctdmFsaWQubmctdG91Y2hlZDpmb2N1cyB7XG4gIGJveC1zaGFkb3c6IDAgMCAwIDAuMnJlbSByZ2JhKDUyLCAxNzksIDEzLCAwLjI1KSAhaW1wb3J0YW50O1xufVxuIl19 */"] });


/***/ }),

/***/ 2136:
/*!**********************************************************!*\
  !*** ./src/app/shopping-list/shopping-list.component.ts ***!
  \**********************************************************/
/***/ ((__unused_webpack_module, __webpack_exports__, __webpack_require__) => {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   "ShoppingListComponent": () => (/* binding */ ShoppingListComponent)
/* harmony export */ });
/* harmony import */ var src_environments_environment__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! src/environments/environment */ 2340);
/* harmony import */ var _angular_core__WEBPACK_IMPORTED_MODULE_4__ = __webpack_require__(/*! @angular/core */ 3184);
/* harmony import */ var _shopping_list_service__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! ./shopping-list.service */ 2457);
/* harmony import */ var _angular_router__WEBPACK_IMPORTED_MODULE_5__ = __webpack_require__(/*! @angular/router */ 2816);
/* harmony import */ var _shared_data_storage_service__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! ../shared/data-storage.service */ 3649);
/* harmony import */ var _angular_common__WEBPACK_IMPORTED_MODULE_6__ = __webpack_require__(/*! @angular/common */ 6362);
/* harmony import */ var _shopping_edit_shopping_edit_component__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! ./shopping-edit/shopping-edit.component */ 51);
/* harmony import */ var _ng_bootstrap_ng_bootstrap__WEBPACK_IMPORTED_MODULE_7__ = __webpack_require__(/*! @ng-bootstrap/ng-bootstrap */ 7544);








function ShoppingListComponent_div_0_ngb_alert_4_Template(rf, ctx) { if (rf & 1) {
    const _r6 = _angular_core__WEBPACK_IMPORTED_MODULE_4__["ɵɵgetCurrentView"]();
    _angular_core__WEBPACK_IMPORTED_MODULE_4__["ɵɵelementStart"](0, "ngb-alert", 9);
    _angular_core__WEBPACK_IMPORTED_MODULE_4__["ɵɵlistener"]("close", function ShoppingListComponent_div_0_ngb_alert_4_Template_ngb_alert_close_0_listener() { _angular_core__WEBPACK_IMPORTED_MODULE_4__["ɵɵrestoreView"](_r6); const ctx_r5 = _angular_core__WEBPACK_IMPORTED_MODULE_4__["ɵɵnextContext"](2); return ctx_r5.ShowMessage = false; });
    _angular_core__WEBPACK_IMPORTED_MODULE_4__["ɵɵtext"](1);
    _angular_core__WEBPACK_IMPORTED_MODULE_4__["ɵɵelementEnd"]();
} if (rf & 2) {
    const ctx_r2 = _angular_core__WEBPACK_IMPORTED_MODULE_4__["ɵɵnextContext"](2);
    _angular_core__WEBPACK_IMPORTED_MODULE_4__["ɵɵproperty"]("type", ctx_r2.MessageType);
    _angular_core__WEBPACK_IMPORTED_MODULE_4__["ɵɵadvance"](1);
    _angular_core__WEBPACK_IMPORTED_MODULE_4__["ɵɵtextInterpolate"](ctx_r2.ResponseFromBackend.Error.Message);
} }
function ShoppingListComponent_div_0_a_6_Template(rf, ctx) { if (rf & 1) {
    const _r9 = _angular_core__WEBPACK_IMPORTED_MODULE_4__["ɵɵgetCurrentView"]();
    _angular_core__WEBPACK_IMPORTED_MODULE_4__["ɵɵelementStart"](0, "a", 10);
    _angular_core__WEBPACK_IMPORTED_MODULE_4__["ɵɵlistener"]("click", function ShoppingListComponent_div_0_a_6_Template_a_click_0_listener() { const restoredCtx = _angular_core__WEBPACK_IMPORTED_MODULE_4__["ɵɵrestoreView"](_r9); const ingredient_r7 = restoredCtx.$implicit; const ctx_r8 = _angular_core__WEBPACK_IMPORTED_MODULE_4__["ɵɵnextContext"](2); return ctx_r8.ShopListServ.SelectItemShopList(ingredient_r7); });
    _angular_core__WEBPACK_IMPORTED_MODULE_4__["ɵɵtext"](1);
    _angular_core__WEBPACK_IMPORTED_MODULE_4__["ɵɵelementStart"](2, "span", 11);
    _angular_core__WEBPACK_IMPORTED_MODULE_4__["ɵɵtext"](3);
    _angular_core__WEBPACK_IMPORTED_MODULE_4__["ɵɵelementEnd"]()();
} if (rf & 2) {
    const ingredient_r7 = ctx.$implicit;
    const ctx_r3 = _angular_core__WEBPACK_IMPORTED_MODULE_4__["ɵɵnextContext"](2);
    _angular_core__WEBPACK_IMPORTED_MODULE_4__["ɵɵproperty"]("ngClass", ctx_r3.ShopListServ.IsCurrentSelected(ingredient_r7) ? "active" : "");
    _angular_core__WEBPACK_IMPORTED_MODULE_4__["ɵɵadvance"](1);
    _angular_core__WEBPACK_IMPORTED_MODULE_4__["ɵɵtextInterpolate1"]("", ingredient_r7.Name, " ");
    _angular_core__WEBPACK_IMPORTED_MODULE_4__["ɵɵadvance"](2);
    _angular_core__WEBPACK_IMPORTED_MODULE_4__["ɵɵtextInterpolate"](ingredient_r7.Amount);
} }
function ShoppingListComponent_div_0_ngb_pagination_7_Template(rf, ctx) { if (rf & 1) {
    const _r11 = _angular_core__WEBPACK_IMPORTED_MODULE_4__["ɵɵgetCurrentView"]();
    _angular_core__WEBPACK_IMPORTED_MODULE_4__["ɵɵelementStart"](0, "ngb-pagination", 12);
    _angular_core__WEBPACK_IMPORTED_MODULE_4__["ɵɵlistener"]("pageChange", function ShoppingListComponent_div_0_ngb_pagination_7_Template_ngb_pagination_pageChange_0_listener($event) { _angular_core__WEBPACK_IMPORTED_MODULE_4__["ɵɵrestoreView"](_r11); const ctx_r10 = _angular_core__WEBPACK_IMPORTED_MODULE_4__["ɵɵnextContext"](2); return ctx_r10.OnPageChanged($event); });
    _angular_core__WEBPACK_IMPORTED_MODULE_4__["ɵɵelementEnd"]();
} if (rf & 2) {
    const ctx_r4 = _angular_core__WEBPACK_IMPORTED_MODULE_4__["ɵɵnextContext"](2);
    _angular_core__WEBPACK_IMPORTED_MODULE_4__["ɵɵproperty"]("collectionSize", ctx_r4.slcollectionSize)("pageSize", ctx_r4.slPageSize)("page", ctx_r4.slCurrentPage)("rotate", true)("boundaryLinks", true)("maxSize", 11);
} }
function ShoppingListComponent_div_0_Template(rf, ctx) { if (rf & 1) {
    _angular_core__WEBPACK_IMPORTED_MODULE_4__["ɵɵelementStart"](0, "div", 3)(1, "div", 4);
    _angular_core__WEBPACK_IMPORTED_MODULE_4__["ɵɵelement"](2, "app-shopping-edit")(3, "hr");
    _angular_core__WEBPACK_IMPORTED_MODULE_4__["ɵɵtemplate"](4, ShoppingListComponent_div_0_ngb_alert_4_Template, 2, 2, "ngb-alert", 5);
    _angular_core__WEBPACK_IMPORTED_MODULE_4__["ɵɵelementStart"](5, "ul", 6);
    _angular_core__WEBPACK_IMPORTED_MODULE_4__["ɵɵtemplate"](6, ShoppingListComponent_div_0_a_6_Template, 4, 3, "a", 7);
    _angular_core__WEBPACK_IMPORTED_MODULE_4__["ɵɵelementEnd"]();
    _angular_core__WEBPACK_IMPORTED_MODULE_4__["ɵɵtemplate"](7, ShoppingListComponent_div_0_ngb_pagination_7_Template, 1, 6, "ngb-pagination", 8);
    _angular_core__WEBPACK_IMPORTED_MODULE_4__["ɵɵelementEnd"]()();
} if (rf & 2) {
    const ctx_r0 = _angular_core__WEBPACK_IMPORTED_MODULE_4__["ɵɵnextContext"]();
    _angular_core__WEBPACK_IMPORTED_MODULE_4__["ɵɵadvance"](4);
    _angular_core__WEBPACK_IMPORTED_MODULE_4__["ɵɵproperty"]("ngIf", ctx_r0.ShowMessage);
    _angular_core__WEBPACK_IMPORTED_MODULE_4__["ɵɵadvance"](2);
    _angular_core__WEBPACK_IMPORTED_MODULE_4__["ɵɵproperty"]("ngForOf", ctx_r0.ingredients);
    _angular_core__WEBPACK_IMPORTED_MODULE_4__["ɵɵadvance"](1);
    _angular_core__WEBPACK_IMPORTED_MODULE_4__["ɵɵproperty"]("ngIf", ctx_r0.slcollectionSize > ctx_r0.slPageSize);
} }
function ShoppingListComponent_div_2_Template(rf, ctx) { if (rf & 1) {
    _angular_core__WEBPACK_IMPORTED_MODULE_4__["ɵɵelementStart"](0, "div", 13)(1, "span", 14);
    _angular_core__WEBPACK_IMPORTED_MODULE_4__["ɵɵtext"](2, "Loading...");
    _angular_core__WEBPACK_IMPORTED_MODULE_4__["ɵɵelementEnd"]()();
} }
class ShoppingListComponent {
    constructor(ShopListServ, activeroute, DataServ, router) {
        this.ShopListServ = ShopListServ;
        this.activeroute = activeroute;
        this.DataServ = DataServ;
        this.router = router;
    }
    ngOnDestroy() {
        this.IngChanged.unsubscribe();
        this.PageChanged.unsubscribe();
        this.FetchOnInint.unsubscribe();
        this.DataLoading.unsubscribe();
        this.RecivedErrorSub.unsubscribe();
        this.WatchIngAdd.unsubscribe();
        this.WatchIngDel.unsubscribe();
        this.WatchIngCle.unsubscribe();
    }
    ngOnInit() {
        this.slPageSize = src_environments_environment__WEBPACK_IMPORTED_MODULE_0__.environment.ShoppingListPageSize;
        this.RecivedErrorSub = this.DataServ.RecivedError.subscribe((response) => {
            this.ShowMessage = true;
            this.ResponseFromBackend = response;
            setTimeout(() => this.ShowMessage = false, 5000);
            if (response) {
                switch (response.Error.Code) {
                    case 200:
                        this.MessageType = 'success';
                        break;
                    default:
                        this.MessageType = 'danger';
                        break;
                }
            }
        });
        this.IngChanged = this.ShopListServ.IngredientChanged
            .subscribe((ing) => {
            this.ingredients = ing;
        });
        this.PageChanged = this.activeroute.params.subscribe((params) => {
            this.slCurrentPage = +params.pn;
        });
        this.DataLoading = this.DataServ.LoadingData.subscribe((State) => {
            this.IsLoading = State;
        });
        this.FetchOnInint = this.DataServ.FetchShoppingList(this.slCurrentPage, src_environments_environment__WEBPACK_IMPORTED_MODULE_0__.environment.ShoppingListPageSize).subscribe((value) => {
            this.ingredients = this.ShopListServ.GetIngredients();
            this.slcollectionSize = this.ShopListServ.Total;
        }, (error) => {
            this.ingredients = [];
        });
        this.WatchIngAdd = this.ShopListServ.IngredientAdded.subscribe((newing) => {
            this.slcollectionSize += 1;
            this.ingredients = this.ShopListServ.GetIngredients();
        });
        this.WatchIngDel = this.ShopListServ.IngredientDeleted.subscribe((deling) => {
            this.slcollectionSize -= 1;
            this.ingredients = this.ShopListServ.GetIngredients();
            if (this.ingredients.length === 0) {
                this.slCurrentPage = this.GetPreviousPage(this.slCurrentPage);
                this.ShopListServ.Total = this.slcollectionSize;
                if (this.slcollectionSize !== 0) {
                    this.OnPageChanged(this.slCurrentPage);
                }
            }
        });
        this.WatchIngCle = this.ShopListServ.IngredientClear.subscribe(() => {
            this.slcollectionSize = 0;
            this.ShopListServ.Total = this.slcollectionSize;
        });
    }
    GetPreviousPage(page) {
        if (page > 1) {
            return page - 1;
        }
        else {
            return 1;
        }
    }
    OnPageChanged(page) {
        this.slCurrentPage = page;
        this.FetchOnInint = this.DataServ.FetchShoppingList(page, src_environments_environment__WEBPACK_IMPORTED_MODULE_0__.environment.ShoppingListPageSize).subscribe(() => {
            this.ingredients = this.ShopListServ.GetIngredients();
            this.router.navigate(['../', page.toString()], { relativeTo: this.activeroute });
        });
    }
}
ShoppingListComponent.ɵfac = function ShoppingListComponent_Factory(t) { return new (t || ShoppingListComponent)(_angular_core__WEBPACK_IMPORTED_MODULE_4__["ɵɵdirectiveInject"](_shopping_list_service__WEBPACK_IMPORTED_MODULE_1__.ShoppingListService), _angular_core__WEBPACK_IMPORTED_MODULE_4__["ɵɵdirectiveInject"](_angular_router__WEBPACK_IMPORTED_MODULE_5__.ActivatedRoute), _angular_core__WEBPACK_IMPORTED_MODULE_4__["ɵɵdirectiveInject"](_shared_data_storage_service__WEBPACK_IMPORTED_MODULE_2__.DataStorageService), _angular_core__WEBPACK_IMPORTED_MODULE_4__["ɵɵdirectiveInject"](_angular_router__WEBPACK_IMPORTED_MODULE_5__.Router)); };
ShoppingListComponent.ɵcmp = /*@__PURE__*/ _angular_core__WEBPACK_IMPORTED_MODULE_4__["ɵɵdefineComponent"]({ type: ShoppingListComponent, selectors: [["app-shopping-list"]], decls: 3, vars: 2, consts: [["class", "row", 4, "ngIf"], [1, "text-center"], ["class", "spinner-border", "role", "status", 4, "ngIf"], [1, "row"], [1, "col"], [3, "type", "close", 4, "ngIf"], [1, "list-group", "mb-1"], ["style", "cursor: pointer;", "class", "list-group-item list-group-item-action d-flex justify-content-between align-items-center", 3, "ngClass", "click", 4, "ngFor", "ngForOf"], [3, "collectionSize", "pageSize", "page", "rotate", "boundaryLinks", "maxSize", "pageChange", 4, "ngIf"], [3, "type", "close"], [1, "list-group-item", "list-group-item-action", "d-flex", "justify-content-between", "align-items-center", 2, "cursor", "pointer", 3, "ngClass", "click"], [1, "badge", "bg-primary", "rounded-pill"], [3, "collectionSize", "pageSize", "page", "rotate", "boundaryLinks", "maxSize", "pageChange"], ["role", "status", 1, "spinner-border"], [1, "visually-hidden-focusable"]], template: function ShoppingListComponent_Template(rf, ctx) { if (rf & 1) {
        _angular_core__WEBPACK_IMPORTED_MODULE_4__["ɵɵtemplate"](0, ShoppingListComponent_div_0_Template, 8, 3, "div", 0);
        _angular_core__WEBPACK_IMPORTED_MODULE_4__["ɵɵelementStart"](1, "div", 1);
        _angular_core__WEBPACK_IMPORTED_MODULE_4__["ɵɵtemplate"](2, ShoppingListComponent_div_2_Template, 3, 0, "div", 2);
        _angular_core__WEBPACK_IMPORTED_MODULE_4__["ɵɵelementEnd"]();
    } if (rf & 2) {
        _angular_core__WEBPACK_IMPORTED_MODULE_4__["ɵɵproperty"]("ngIf", !ctx.IsLoading);
        _angular_core__WEBPACK_IMPORTED_MODULE_4__["ɵɵadvance"](2);
        _angular_core__WEBPACK_IMPORTED_MODULE_4__["ɵɵproperty"]("ngIf", ctx.IsLoading);
    } }, directives: [_angular_common__WEBPACK_IMPORTED_MODULE_6__.NgIf, _shopping_edit_shopping_edit_component__WEBPACK_IMPORTED_MODULE_3__.ShoppingEditComponent, _ng_bootstrap_ng_bootstrap__WEBPACK_IMPORTED_MODULE_7__.NgbAlert, _angular_common__WEBPACK_IMPORTED_MODULE_6__.NgForOf, _angular_common__WEBPACK_IMPORTED_MODULE_6__.NgClass, _ng_bootstrap_ng_bootstrap__WEBPACK_IMPORTED_MODULE_7__.NgbPagination], styles: ["\n/*# sourceMappingURL=data:application/json;base64,eyJ2ZXJzaW9uIjozLCJzb3VyY2VzIjpbXSwibmFtZXMiOltdLCJtYXBwaW5ncyI6IiIsImZpbGUiOiJzaG9wcGluZy1saXN0LmNvbXBvbmVudC5jc3MifQ== */"] });


/***/ }),

/***/ 6673:
/*!*******************************************************!*\
  !*** ./src/app/shopping-list/shopping-list.module.ts ***!
  \*******************************************************/
/***/ ((__unused_webpack_module, __webpack_exports__, __webpack_require__) => {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   "ShoppingListModule": () => (/* binding */ ShoppingListModule)
/* harmony export */ });
/* harmony import */ var _angular_common__WEBPACK_IMPORTED_MODULE_4__ = __webpack_require__(/*! @angular/common */ 6362);
/* harmony import */ var _shopping_list_component__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ./shopping-list.component */ 2136);
/* harmony import */ var _shopping_edit_shopping_edit_component__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! ./shopping-edit/shopping-edit.component */ 51);
/* harmony import */ var _angular_router__WEBPACK_IMPORTED_MODULE_7__ = __webpack_require__(/*! @angular/router */ 2816);
/* harmony import */ var _angular_forms__WEBPACK_IMPORTED_MODULE_5__ = __webpack_require__(/*! @angular/forms */ 587);
/* harmony import */ var _ng_bootstrap_ng_bootstrap__WEBPACK_IMPORTED_MODULE_6__ = __webpack_require__(/*! @ng-bootstrap/ng-bootstrap */ 7544);
/* harmony import */ var _auth_role_guard__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! ../auth/role.guard */ 9805);
/* harmony import */ var _angular_core__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! @angular/core */ 3184);









const routes = [
    { path: '', redirectTo: '1', pathMatch: 'full' },
    { path: ':pn', component: _shopping_list_component__WEBPACK_IMPORTED_MODULE_0__.ShoppingListComponent, canActivate: [_auth_role_guard__WEBPACK_IMPORTED_MODULE_2__.RoleGuard], data: { expectedRole: 'admin_role_CRUD' } }
];
class ShoppingListModule {
}
ShoppingListModule.ɵfac = function ShoppingListModule_Factory(t) { return new (t || ShoppingListModule)(); };
ShoppingListModule.ɵmod = /*@__PURE__*/ _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵdefineNgModule"]({ type: ShoppingListModule });
ShoppingListModule.ɵinj = /*@__PURE__*/ _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵdefineInjector"]({ imports: [[
            _angular_common__WEBPACK_IMPORTED_MODULE_4__.CommonModule,
            _angular_forms__WEBPACK_IMPORTED_MODULE_5__.FormsModule,
            _ng_bootstrap_ng_bootstrap__WEBPACK_IMPORTED_MODULE_6__.NgbAlertModule,
            _ng_bootstrap_ng_bootstrap__WEBPACK_IMPORTED_MODULE_6__.NgbPaginationModule,
            _angular_router__WEBPACK_IMPORTED_MODULE_7__.RouterModule.forChild(routes)
        ]] });
(function () { (typeof ngJitMode === "undefined" || ngJitMode) && _angular_core__WEBPACK_IMPORTED_MODULE_3__["ɵɵsetNgModuleScope"](ShoppingListModule, { declarations: [_shopping_list_component__WEBPACK_IMPORTED_MODULE_0__.ShoppingListComponent,
        _shopping_edit_shopping_edit_component__WEBPACK_IMPORTED_MODULE_1__.ShoppingEditComponent], imports: [_angular_common__WEBPACK_IMPORTED_MODULE_4__.CommonModule,
        _angular_forms__WEBPACK_IMPORTED_MODULE_5__.FormsModule,
        _ng_bootstrap_ng_bootstrap__WEBPACK_IMPORTED_MODULE_6__.NgbAlertModule,
        _ng_bootstrap_ng_bootstrap__WEBPACK_IMPORTED_MODULE_6__.NgbPaginationModule, _angular_router__WEBPACK_IMPORTED_MODULE_7__.RouterModule] }); })();


/***/ })

}]);
//# sourceMappingURL=src_app_shopping-list_shopping-list_module_ts.js.map