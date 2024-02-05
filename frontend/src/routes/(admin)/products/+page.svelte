<script>
    // @ts-nocheck
    
    import { onMount } from "svelte";
    import Icon from "@iconify/svelte";
  import AdminNav from "../../../components/adminNav.svelte";
    let products = []

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
</script>

<div class="w-screen h-screen flex">
    <AdminNav />
    <div class="w-full flex flex-col m-14 h-screen">
        <div class="flex flex-wrap justify-center items-center space-x-10 gap-y-10 pt-10">
            {#each products as product(product.ID)}
            <a href={`get/${product.ID}`} class="bg-gray-100 h-64 w-48 rounded-2xl flex flex-col space-y-2 items-center  shadow-xl">
                <div class="flex h-52 bg-white w-full rounded-2xl">
                    <img src={`http://localhost:8000/${product.image1}`}  alt="" srcset="" class="object-cover h-full w-full rounded-t-2xl">
                </div>                    
                <p class="text-teal-900 font-semibold">{product.name}</p>
            </a>       
            {/each}                              
        </div>         
    </div>
     
</div>
