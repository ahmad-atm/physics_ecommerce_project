<script>
    import Icon from "@iconify/svelte";
    import { onMount } from "svelte";
    
    onMount(()=>{
         fetch("http://localhost:8000/profile", {
            headers:{
                "content-type":"application/json",
                authorization:`Bearer ${localStorage.getItem('ecommerce_token')}`
            }
        })
        .then(res=>res.json())
        .then(data=>{user=data; console.log(data)})
        .catch(err=>console.log(err))
    })
    let user = {}
</script>

<div class="h-screen bg-slate-900 w-[12%] flex flex-col">
    <div class="h-[70%] flex items-center justify-center">
        <div class="flex flex-col items-center space-y-1">
            <div class="bg-white h-24 w-24 rounded-full flex items-center justify-center">
                <img src={`http://localhost:8000/static/${user.profileImage}`} alt="" class="h-full w-full rounded-full">
            </div>
            <p class="text-gray-200 text-sm font-bold">{user.name}</p>
        </div>
    </div>
    <div class="bg-slate-800 rounded-tr-[2.5rem] h-full flex flex-col items-center justify-center space-y-3">
        <a href="/dashboard" class="bg-gray-400 hover:bg-slate-900 rounded-full px-4 py-4 text-slate-900 hover:text-slate-400">
            <Icon icon="ic:round-dashboard"  class="text-xl"/>
        </a>
        <a href="/products" class="bg-gray-400 hover:bg-slate-900 rounded-full px-4 py-4 text-slate-900 hover:text-slate-400">
            <Icon icon="dashicons:products"  class="text-xl"/>
        </a>
        <a href="/orders" class="bg-gray-400 hover:bg-slate-900 rounded-full px-4 py-4 text-slate-900 hover:text-slate-400">
            <Icon icon="material-symbols:draft-orders-sharp"  class="text-xl"/>
        </a>       
        <a href="/logout" class="bg-gray-400 hover:bg-slate-900 rounded-full px-4 py-4 text-slate-900 hover:text-slate-400">
            <Icon icon="fe:logout" class="text-xl"/>
        </a>             
    </div>
</div>
