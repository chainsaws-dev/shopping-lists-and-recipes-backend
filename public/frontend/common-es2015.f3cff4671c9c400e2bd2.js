(window.webpackJsonp=window.webpackJsonp||[]).push([[1],{EiSp:function(e,t,i){"use strict";i.d(t,"a",(function(){return n}));class n{constructor(e,t){this.Name=e,this.Amount=t}}},GXvH:function(e,t,i){"use strict";i.d(t,"a",(function(){return p}));var n=i("tk/3"),s=i("lJxs"),r=i("vkgz"),c=i("AytR"),d=i("XNiG"),h=i("fXoL"),a=i("ceC1"),o=i("ozzT");let p=(()=>{class e{constructor(e,t,i){this.http=e,this.recipes=t,this.shoppinglist=i,this.LoadingData=new d.a,this.RecivedResponse=new d.a}DeleteRecipe(e){this.LoadingData.next(!0);const t={headers:new n.d({RecipeID:e.ID.toString()})};this.http.delete(c.a.GetSetRecipesUrl,t).subscribe(e=>{this.RecivedResponse.next(e),this.LoadingData.next(!1)})}SaveRecipe(e){this.LoadingData.next(!0),this.http.post(c.a.GetSetRecipesUrl,e).subscribe(e=>{this.RecivedResponse.next(e),this.LoadingData.next(!1)})}FetchRecipes(e,t){this.LoadingData.next(!0);const i={headers:new n.d({Page:e.toString(),Limit:t.toString()})};return this.http.get(c.a.GetSetRecipesUrl,i).pipe(Object(s.a)(e=>(e.Recipes=e.Recipes.map(e=>Object.assign(Object.assign({},e),{Ingredients:e.Ingredients?e.Ingredients:[]})),e)),Object(r.a)(e=>{this.recipes.SetRecipes(e.Recipes),this.recipes.SetPagination(e.Total,e.Limit,e.Offset),this.LoadingData.next(!1)}))}SearchRecipes(e,t,i){this.LoadingData.next(!0);const d={headers:new n.d({Page:e.toString(),Limit:t.toString(),Search:i})};return this.http.get(c.a.SearchRecipesUrl,d).pipe(Object(s.a)(e=>(e.Recipes=e.Recipes.map(e=>Object.assign(Object.assign({},e),{Ingredients:e.Ingredients?e.Ingredients:[]})),e)),Object(r.a)(e=>{this.recipes.SetRecipes(e.Recipes),this.recipes.SetPagination(e.Total,e.Limit,e.Offset),this.LoadingData.next(!1)}))}FetchShoppingList(e,t){this.LoadingData.next(!0);const i={headers:new n.d({Page:e.toString(),Limit:t.toString()})};return this.http.get(c.a.GetSetShoppingListUrl,i).pipe(Object(r.a)(e=>{this.shoppinglist.SetIngredients(e.Items),this.shoppinglist.SetPagination(e.Total,e.Limit,e.Offset),this.LoadingData.next(!1)}))}}return e.\u0275fac=function(t){return new(t||e)(h.Xb(n.a),h.Xb(a.a),h.Xb(o.a))},e.\u0275prov=h.Gb({token:e,factory:e.\u0275fac,providedIn:"root"}),e})()},ceC1:function(e,t,i){"use strict";i.d(t,"a",(function(){return a}));var n=i("qc5x"),s=i("EiSp"),r=i("XNiG"),c=i("AytR"),d=i("fXoL"),h=i("ozzT");let a=(()=>{class e{constructor(e){this.ShopList=e,this.IngredientSelected=new r.a,this.RecipeChanged=new r.a,this.RecipesUpdated=new r.a,this.RecipesInserted=new r.a,this.RecipesDeleted=new r.a,this.recipes=[]}GetRecipes(){return this.recipes.slice(0,c.a.RecipePageSize)}GetRecipesLen(){return this.recipes.length}GetRecipeById(e){return e<this.recipes.length&&e>0?this.recipes[e]:this.recipes[0]}GetRecipeId(e){return this.recipes.indexOf(e)}SendToShoppingList(e){e.forEach(e=>{this.ShopList.AddNewItem(e)})}AddNewRecipe(e){const t=new n.a(e.Name,e.Description,e.ImagePath,e.Ingredients);this.recipes.length<=c.a.RecipePageSize&&this.recipes.push(t),this.RecipeChanged.next(t),this.RecipesInserted.next()}UpdateExistingRecipe(e,t){this.recipes[t]=e,this.RecipeChanged.next(e)}SetRecipes(e){this.recipes=e,this.RecipesUpdated.next()}SetPagination(e,t,i){this.Total=e}DeleteRecipe(e){this.recipes.splice(e,1),this.RecipesDeleted.next()}AddNewIngredient(e){const t=this.RecipeToEdit.Ingredients.find(t=>t.Name===e.Name);t?t.Amount=t.Amount+e.Amount:this.RecipeToEdit.Ingredients.push(new s.a(e.Name,e.Amount))}UpdateSelectedIngredient(e){const t=this.RecipeToEdit.Ingredients.find(e=>e===this.CurrentSelectedItem);t&&(t.Name=e.Name,t.Amount=e.Amount),this.CurrentSelectedItem=null,this.IngredientSelected.next(this.CurrentSelectedItem)}DeleteSelectedIngredient(){const e=this.RecipeToEdit.Ingredients.indexOf(this.CurrentSelectedItem);-1!==e&&this.RecipeToEdit.Ingredients.splice(e,1),this.CurrentSelectedItem=null}ClearAllIngredients(){this.RecipeToEdit.Ingredients=[],this.CurrentSelectedItem=null}GetIngredientsLength(){return this.RecipeToEdit.Ingredients.length}IngredientSelect(e){this.CurrentSelectedItem=e,this.IngredientSelected.next(e)}}return e.\u0275fac=function(t){return new(t||e)(d.Xb(h.a))},e.\u0275prov=d.Gb({token:e,factory:e.\u0275fac,providedIn:"root"}),e})()},ozzT:function(e,t,i){"use strict";i.d(t,"a",(function(){return d}));var n=i("EiSp"),s=i("XNiG"),r=i("AytR"),c=i("fXoL");let d=(()=>{class e{constructor(){this.IngredientSelected=new s.a,this.IngredientChanged=new s.a,this.IngredientsUpdated=new s.a,this.ingredients=[]}GetIngredients(){return this.ingredients.slice()}GetIngredientsLength(){return this.ingredients.length}SetIngredients(e){this.ingredients=e,this.IngredientsUpdated.next()}SetPagination(e,t,i){this.Total=e}AddNewItem(e){const t=this.ingredients.find(t=>t.Name===e.Name);t?t.Amount=t.Amount+e.Amount:this.ingredients.length<=r.a.ShoppingListPageSize&&this.ingredients.push(new n.a(e.Name,e.Amount)),this.IngredientChanged.next(this.ingredients.slice())}UpdateSelectedItem(e){const t=this.ingredients.indexOf(this.CurrentSelectedItem);-1!==t&&(this.ingredients[t]=e,this.IngredientChanged.next(this.ingredients.slice())),this.CurrentSelectedItem=null}DeleteSelectedItem(){const e=this.ingredients.indexOf(this.CurrentSelectedItem);-1!==e&&this.ingredients.splice(e,1),this.CurrentSelectedItem=null,this.IngredientChanged.next(this.ingredients.slice())}ClearAll(){this.ingredients=[],this.CurrentSelectedItem=null,this.IngredientChanged.next(this.ingredients.slice())}SelectItemShopList(e){this.CurrentSelectedItem=e,this.IngredientSelected.next(e)}IsCurrentSelected(e){return this.CurrentSelectedItem===e}}return e.\u0275fac=function(t){return new(t||e)},e.\u0275prov=c.Gb({token:e,factory:e.\u0275fac,providedIn:"root"}),e})()},qc5x:function(e,t,i){"use strict";i.d(t,"a",(function(){return n}));class n{constructor(e,t,i,n){this.Name=e,this.Description=t,this.ImagePath=i,this.Ingredients=n}}}}]);