<script>
    import Nav from "../../../../components/nav.svelte";
    import { onMount } from "svelte";
    import { page } from "$app/stores";

    let product = {}
    const id = $page.params.productId
    let msg = "Added to Cart"
    let cart = "Add to Cart"
    let cartText = cart
    
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
          console.log(data)
        })
        .catch(err=>console.log(err))
    }) 

    function addCart(id, name, image, price){
        fetch("http://localhost:8000/add_cart", {
            method:"Post",
            headers:{
                "content-type":"application/json",
                authorization:`Bearer ${localStorage.getItem('ecommerce_token')}`
            },
            body:JSON.stringify({
                productId:id,     
                productName:name, 
                productImage:image,
                productPrice:price,
                quantity: "1"
            })
        })
        .then(res=>res.json())
        .then(data=>{console.log(data); cartText=msg})
        .catch(err=>console.log(err))
    }
</script>

<div class="bg-white h-screen p-1">
    <div class="bg-yellow-200 h-screen mx-[0.02rem] rounded-t-2xl mt-[0.02rem] p-5">
        <Nav />
        <div class="p-2 mt-10 mx-16 h-96 bg-yellow-300 rounded-3xl grid grid-cols-3">
            <div class="flex flex-col space-y-3 justify-center items-center">
                <p class="text-indigo-800 font-bold text-3xl flex flex-wrap">{product.name}</p>
                <p class="text-indigo-800 font-bold text-3xl flex flex-wrap">N{product.price}</p>
                <p class="text-indigo-800 font-medium flex flex-wrap items-center justify-center">{product.description}</p>
            </div>
            <div class="flex items-center justify-center">
                <!-- <p>2</p> -->
                <img src={`http://localhost:8000/${product.image1}`}  alt="" class="h-80 w-80 rounded-full object-cover">
            </div>
            <div class="flex flex-col space-y-3 justify-center items-center">
                <p class="text-indigo-800 font-bold text-3xl flex flex-wrap">Review: {product.name}</p>
                <p class="text-indigo-800 font-medium flex flex-wrap items-center justify-center">Color: {product.description}</p>
                <button on:click={()=> addCart(product.ID, product.name, product.image1, product.price)} 
                    class="px-2 py-5 bg-black rounded-3xl mx-10 text-yellow-300 font-bold w-full hover:bg-gray-800">{cartText}</button>             
            </div>
         </div>
    </div>
</div>