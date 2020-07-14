function _defineProperties(e,t){for(var n=0;n<t.length;n++){var i=t[n];i.enumerable=i.enumerable||!1,i.configurable=!0,"value"in i&&(i.writable=!0),Object.defineProperty(e,i.key,i)}}function _createClass(e,t,n){return t&&_defineProperties(e.prototype,t),n&&_defineProperties(e,n),e}function _classCallCheck(e,t){if(!(e instanceof t))throw new TypeError("Cannot call a class as a function")}(window.webpackJsonp=window.webpackJsonp||[]).push([[1],{EiSp:function(e,t,n){"use strict";n.d(t,"a",(function(){return i}));var i=function e(t,n){_classCallCheck(this,e),this.Name=t,this.Amount=n}},GXvH:function(e,t,n){"use strict";n.d(t,"a",(function(){return h}));var i=n("tk/3"),s=n("lJxs"),r=n("vkgz"),a=n("AytR"),c=n("XNiG"),u=n("fXoL"),o=n("ceC1"),d=n("ozzT"),h=function(){var e=function(){function e(t,n,i){_classCallCheck(this,e),this.http=t,this.recipes=n,this.shoppinglist=i,this.LoadingData=new c.a,this.RecivedResponse=new c.a}return _createClass(e,[{key:"DeleteRecipe",value:function(e){var t=this;this.LoadingData.next(!0);var n={headers:new i.d({RecipeID:e.ID.toString()})};this.http.delete(a.a.GetSetRecipesUrl,n).subscribe((function(e){t.RecivedResponse.next(e),t.LoadingData.next(!1)}))}},{key:"SaveRecipe",value:function(e){var t=this;this.LoadingData.next(!0),this.http.post(a.a.GetSetRecipesUrl,e).subscribe((function(e){t.RecivedResponse.next(e),t.LoadingData.next(!1)}))}},{key:"FetchRecipes",value:function(e,t){var n=this;this.LoadingData.next(!0);var c={headers:new i.d({Page:e.toString(),Limit:t.toString()})};return this.http.get(a.a.GetSetRecipesUrl,c).pipe(Object(s.a)((function(e){return e.Recipes=e.Recipes.map((function(e){return Object.assign(Object.assign({},e),{Ingredients:e.Ingredients?e.Ingredients:[]})})),e})),Object(r.a)((function(e){n.recipes.SetRecipes(e.Recipes),n.recipes.SetPagination(e.Total,e.Limit,e.Offset),n.LoadingData.next(!1)})))}},{key:"SearchRecipes",value:function(e,t,n){var c=this;this.LoadingData.next(!0);var u={headers:new i.d({Page:e.toString(),Limit:t.toString(),Search:n})};return this.http.get(a.a.SearchRecipesUrl,u).pipe(Object(s.a)((function(e){return e.Recipes=e.Recipes.map((function(e){return Object.assign(Object.assign({},e),{Ingredients:e.Ingredients?e.Ingredients:[]})})),e})),Object(r.a)((function(e){c.recipes.SetRecipes(e.Recipes),c.recipes.SetPagination(e.Total,e.Limit,e.Offset),c.LoadingData.next(!1)})))}},{key:"FetchShoppingList",value:function(e,t){var n=this;this.LoadingData.next(!0);var s={headers:new i.d({Page:e.toString(),Limit:t.toString()})};return this.http.get(a.a.GetSetShoppingListUrl,s).pipe(Object(r.a)((function(e){n.shoppinglist.SetIngredients(e.Items),n.shoppinglist.SetPagination(e.Total,e.Limit,e.Offset),n.LoadingData.next(!1)})))}}]),e}();return e.\u0275fac=function(t){return new(t||e)(u.Xb(i.a),u.Xb(o.a),u.Xb(d.a))},e.\u0275prov=u.Gb({token:e,factory:e.\u0275fac,providedIn:"root"}),e}()},ceC1:function(e,t,n){"use strict";n.d(t,"a",(function(){return o}));var i=n("qc5x"),s=n("EiSp"),r=n("XNiG"),a=n("AytR"),c=n("fXoL"),u=n("ozzT"),o=function(){var e=function(){function e(t){_classCallCheck(this,e),this.ShopList=t,this.IngredientSelected=new r.a,this.RecipeChanged=new r.a,this.RecipesUpdated=new r.a,this.RecipesInserted=new r.a,this.RecipesDeleted=new r.a,this.recipes=[]}return _createClass(e,[{key:"GetRecipes",value:function(){return this.recipes.slice(0,a.a.RecipePageSize)}},{key:"GetRecipesLen",value:function(){return this.recipes.length}},{key:"GetRecipeById",value:function(e){return e<this.recipes.length&&e>0?this.recipes[e]:this.recipes[0]}},{key:"GetRecipeId",value:function(e){return this.recipes.indexOf(e)}},{key:"SendToShoppingList",value:function(e){var t=this;e.forEach((function(e){t.ShopList.AddNewItem(e)}))}},{key:"AddNewRecipe",value:function(e){var t=new i.a(e.Name,e.Description,e.ImagePath,e.Ingredients);this.recipes.length<=a.a.RecipePageSize&&this.recipes.push(t),this.RecipeChanged.next(t),this.RecipesInserted.next()}},{key:"UpdateExistingRecipe",value:function(e,t){this.recipes[t]=e,this.RecipeChanged.next(e)}},{key:"SetRecipes",value:function(e){this.recipes=e,this.RecipesUpdated.next()}},{key:"SetPagination",value:function(e,t,n){this.Total=e}},{key:"DeleteRecipe",value:function(e){this.recipes.splice(e,1),this.RecipesDeleted.next()}},{key:"AddNewIngredient",value:function(e){var t=this.RecipeToEdit.Ingredients.find((function(t){return t.Name===e.Name}));t?t.Amount=t.Amount+e.Amount:this.RecipeToEdit.Ingredients.push(new s.a(e.Name,e.Amount))}},{key:"UpdateSelectedIngredient",value:function(e){var t=this,n=this.RecipeToEdit.Ingredients.find((function(e){return e===t.CurrentSelectedItem}));n&&(n.Name=e.Name,n.Amount=e.Amount),this.CurrentSelectedItem=null,this.IngredientSelected.next(this.CurrentSelectedItem)}},{key:"DeleteSelectedIngredient",value:function(){var e=this.RecipeToEdit.Ingredients.indexOf(this.CurrentSelectedItem);-1!==e&&this.RecipeToEdit.Ingredients.splice(e,1),this.CurrentSelectedItem=null}},{key:"ClearAllIngredients",value:function(){this.RecipeToEdit.Ingredients=[],this.CurrentSelectedItem=null}},{key:"GetIngredientsLength",value:function(){return this.RecipeToEdit.Ingredients.length}},{key:"IngredientSelect",value:function(e){this.CurrentSelectedItem=e,this.IngredientSelected.next(e)}}]),e}();return e.\u0275fac=function(t){return new(t||e)(c.Xb(u.a))},e.\u0275prov=c.Gb({token:e,factory:e.\u0275fac,providedIn:"root"}),e}()},ozzT:function(e,t,n){"use strict";n.d(t,"a",(function(){return c}));var i=n("EiSp"),s=n("XNiG"),r=n("AytR"),a=n("fXoL"),c=function(){var e=function(){function e(){_classCallCheck(this,e),this.IngredientSelected=new s.a,this.IngredientChanged=new s.a,this.IngredientsUpdated=new s.a,this.ingredients=[]}return _createClass(e,[{key:"GetIngredients",value:function(){return this.ingredients.slice()}},{key:"GetIngredientsLength",value:function(){return this.ingredients.length}},{key:"SetIngredients",value:function(e){this.ingredients=e,this.IngredientsUpdated.next()}},{key:"SetPagination",value:function(e,t,n){this.Total=e}},{key:"AddNewItem",value:function(e){var t=this.ingredients.find((function(t){return t.Name===e.Name}));t?t.Amount=t.Amount+e.Amount:this.ingredients.length<=r.a.ShoppingListPageSize&&this.ingredients.push(new i.a(e.Name,e.Amount)),this.IngredientChanged.next(this.ingredients.slice())}},{key:"UpdateSelectedItem",value:function(e){var t=this.ingredients.indexOf(this.CurrentSelectedItem);-1!==t&&(this.ingredients[t]=e,this.IngredientChanged.next(this.ingredients.slice())),this.CurrentSelectedItem=null}},{key:"DeleteSelectedItem",value:function(){var e=this.ingredients.indexOf(this.CurrentSelectedItem);-1!==e&&this.ingredients.splice(e,1),this.CurrentSelectedItem=null,this.IngredientChanged.next(this.ingredients.slice())}},{key:"ClearAll",value:function(){this.ingredients=[],this.CurrentSelectedItem=null,this.IngredientChanged.next(this.ingredients.slice())}},{key:"SelectItemShopList",value:function(e){this.CurrentSelectedItem=e,this.IngredientSelected.next(e)}},{key:"IsCurrentSelected",value:function(e){return this.CurrentSelectedItem===e}}]),e}();return e.\u0275fac=function(t){return new(t||e)},e.\u0275prov=a.Gb({token:e,factory:e.\u0275fac,providedIn:"root"}),e}()},qc5x:function(e,t,n){"use strict";n.d(t,"a",(function(){return i}));var i=function e(t,n,i,s){_classCallCheck(this,e),this.Name=t,this.Description=n,this.ImagePath=i,this.Ingredients=s}}}]);