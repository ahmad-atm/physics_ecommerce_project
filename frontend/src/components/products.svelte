<script>
// @ts-nocheck

    import { onMount } from "svelte";

    let products = []
    let msg = "Added to Cart"
    let cart = "Cart"
    let cartText = cart

    onMount(()=>{
        fetch("http://localhost:8000/products", {
            headers:{
                "content-type":"application/json",
                authorization:`Bearer ${localStorage.getItem("ecommerce_token")}`
            },
        })
        .then(res=>res.json())
        .then(data=>{products=data; console.log(data)})
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
                productPrice:price
            })
        })
        .then(res=>res.json())
        .then(data=>{console.log(data); cartText=msg})
        .catch(err=>console.log(err))
    }
</script>

<div class="flex flex-wrap justify-center items-center space-x-10 gap-y-10 pt-10">
    {#each products as product(product.ID)}
    <div  class="bg-gray-100 h-[17rem] w-48 rounded-2xl flex flex-col space-y-2 items-center  shadow-xl">
        <a href={`shop/${product.ID}`} class="flex h-52 bg-white w-full rounded-2xl">
            <img src={`http://localhost:8000/${product.image1}`}  alt="" srcset="" class="object-cover h-52 w-full rounded-t-2xl">
        </a>   
        <div class="flex flex-col  w-full items-center">
            <p class="text-indigo-600 font-semibold">{product.name}</p>
            <p class="text-gray-800 font-semibold">N{product.price}</p>
            <!-- <button on:click={()=> addCart(product.ID, product.name, product.image1, product.price)} class="p-2 bg-orange-500 rounded-md text-white font-bold w-full hover:bg-orange-300">{cartText}</button>              -->
        </div>                 
    </div>          
    {/each}                              
</div>