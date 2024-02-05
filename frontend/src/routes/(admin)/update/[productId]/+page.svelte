<script>
  // @ts-nocheck
  import { goto } from "$app/navigation";
  import AdminNav from "../../../../components/adminNav.svelte";
  import { page } from "$app/stores";
  import { onMount } from "svelte";

  const id = $page.params.productId
  let product = {}
  let name = "";
  let price = ""; 
  let description = "";
  let category = ""
 
  onMount(()=>{
    fetch(`http://localhost:8000/get_product/${id}`,{
      headers:{
                "content-type":"application/json",
                authorization:`Bearer ${localStorage.getItem('ecommerce_token')}`
            }
        })
        .then(res=>res.json())
        .then(data=>{
          product=data
          name = product.name
          price = product.price
          description = product.description
          category = product.category
        })
        .catch(err=>console.log(err))
    }) 
  
    function updateProduct() {
      var formData =  new FormData()
      formData.append("name", name)
      formData.append("price", price)
      formData.append("description", description)
      formData.append("category", category)
      formData.append("productId", product.ID)
      // @ts-ignore
      fetch("http://localhost:8000/update_product",{
        method:"Post",
        headers:{
          "accept":"*/*",
          authorization:`Bearer ${localStorage.getItem('ecommerce_token')}`
        },
        body:formData
      })
      .then(res=>res.json())
      .then(data=>{console.log(data); goto("/dashboard")})
      .catch(err=>console.log(err))
      
    }
  </script>
  
  <div class="flex h-screen">
      <AdminNav />
      <div class="m-14 p-4 bg-white shadow-xl w-full shadow-gray-400 rounded-lg flex flex-col items-center">
          <p class="flex items-center justify-center text-gray-800 font-semibold">Update {product.name} </p>
          <form enctype="multipart/form-data" class="pt-5 flex flex-col space-y-10 justify-center" on:submit|preventDefault={updateProduct}>
            <div class="flex space-x-14">
              <input type="text" bind:value={name} placeholder="Product Name" class="rounded-xl bg-gray-200 p-3 outline-none">
              <input type="text" bind:value={price} placeholder="Product Price NGN" class="rounded-xl bg-gray-200 p-3 outline-none">
            </div>
  
            <div class="flex justify-center">
              <input type="text" bind:value={description}  placeholder="Product Description" class="rounded-lg bg-gray-200 p-3 outline-none w-full h-40">  
            </div>
  
            <div class="flex flex-col space-y-4">
              <div>
                <label for="category">Category</label>
                <select name="category" id="category" bind:value={category} required>
                  <option value="none" selected disabled hidden>Select an Option</option> 
                  <option value="Laptop">Laptop</option>
                  <option value="Console">Console</option>
                  <option value="Phone">Phone</option>
                  <option value="Phone Accessory">Phone Accessory</option>
                  <option value="Watch">Watch</option>
                  <option value="Wallpaper">Wallpaper</option>
                  <option value="Other">Other</option>
                </select>
              </div>
            </div>
            <div class="flex flex-col space-y-4">
              <button type="submit" class="bg-blue-600 p-3 rounded-md text-white">Update Product</button>
            </div>
  
          </form>
      </div>
  </div>
  