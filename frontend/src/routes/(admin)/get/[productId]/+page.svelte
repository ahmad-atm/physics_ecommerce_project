<script>
    import AdminNav from "../../../../components/adminNav.svelte";
    import { onMount } from "svelte";
    import { page } from "$app/stores";

    let product = {}
    const id = $page.params.productId
    
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
        })
        .catch(err=>console.log(err))
    }) 
</script>

<div class="w-screen h-screen flex">
    <AdminNav />    
</div>
