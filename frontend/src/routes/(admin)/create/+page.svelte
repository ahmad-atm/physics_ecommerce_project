<script>
  import { goto } from "$app/navigation";
  import AdminNav from "../../../components/adminNav.svelte";

// @ts-nocheck


  import Nav from "../../../components/adminNav.svelte";
  let name = "";
  let price = ""; 
  let description = "";
  let category = ""
  let files = [];
  // let file1;
  // let file2=[];
  // let file3=[];


  function createProduct() {
    var formData =  new FormData()
    formData.append("name", name)
    formData.append("price", price)
    formData.append("description", description)
    formData.append("category", category)
    // @ts-ignore
    formData.append("image1", files[0])
    // @ts-ignore
    // formData.append("image1", files[1])
    // // @ts-ignore
    // formData.append("image2", files2)
    fetch("http://localhost:8000/create_product",{
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
        <p class="flex items-center justify-center text-gray-800 font-semibold">Add New Product </p>
        <form enctype="multipart/form-data" class="pt-5 flex flex-col space-y-10 justify-center" on:submit|preventDefault={createProduct}>
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
            <div>
              <label for="image1">Image 1</label>
              <input type="file" bind:files id="image1" accept=".png, .jpg, .jpeg" required >              
            </div>
            <button type="submit" class="bg-blue-600 p-3 rounded-md text-white">Add New  Product</button>
          </div>

        </form>
    </div>
</div>
