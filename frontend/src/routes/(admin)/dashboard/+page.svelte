<script>
// @ts-nocheck

    import Icon from '@iconify/svelte';
    import AdminNav from '../../../components/adminNav.svelte';
    import {onMount} from "svelte"

    let latestProduct = []
    let count = {}

    onMount(async()=>{
        await fetch("http://localhost:8000/latest_products", {
            headers:{
                "content-type":"application/json",
                authorization:`Bearer ${localStorage.getItem('ecommerce_token')}`
            }
        })
        .then(res=>res.json())
        .then(data=>latestProduct=data)
        .catch(err=>console.log(err))

        await fetch("http://localhost:8000/count", {
            headers:{
                "content-type":"application/json",
                authorization:`Bearer ${localStorage.getItem('ecommerce_token')}`
            }
        })
        .then(res=>res.json())
        .then(data=>count=data)
        .catch(err=>console.log(err))
    })    

</script>


<div class="w-screen h-screen flex">
    <!-- <Nav /> -->
    <AdminNav />

    <div class="w-full flex flex-col m-14">
        <div class="flex flex-col space-y-7">
            <div class="flex space-x-4  items-center">
                <a href="/dashboard" class="bg-teal-400 h-60 w-48 rounded-2xl p-5 flex flex-col items-center justify-center shadow-xl">                    
                    <Icon icon="ph:users-bold"  class="text-4xl text-teal-800"/>
                    <p class="text-teal-900 font-semibold">{count.userCount}</p>
                </a>
                <a href="/products" class="bg-indigo-500 h-60 w-48 rounded-2xl p-5 flex flex-col items-center justify-center shadow-xl">                    
                    <Icon icon="dashicons:products"  class="text-4xl text-indigo-900"/>
                    <p class="text-indigo-900 font-semibold">{count.productCount}</p>
                </a>     
                <a href="/orders" class="bg-rose-400 h-60 w-48  rounded-2xl p-5 flex flex-col items-center justify-center shadow-xl">                    
                    <Icon icon="material-symbols:draft-orders-sharp"  class="text-4xl text-rose-800"/>
                    <p class="text-rose-900 font-semibold">{count.orderCount}</p>
                </a> 
                <a href="/create" class="bg-gray-300 h-60 w-48  rounded-2xl p-5 flex flex-col items-center justify-center shadow-xl">                    
                    <Icon icon="typcn:plus"  class="text-4xl text-gray-700"/>
                    <!-- <p class="text-rose-900 font-semibold">02</p> -->
                </a>                             
            </div>    
            
            <div class="bg-white w-full h-full shadow-md">
                <div class="relative overflow-x-auto shadow-md sm:rounded-lg">
                    <table class="w-full text-sm text-left rtl:text-right text-gray-500 ">
                        <caption class="p-5 text-md font-semibold text-left rtl:text-right text-gray-900 bg-white ">
                            Latest products
                        </caption>
                        <thead class="text-xs text-gray-700 uppercase bg-gray-50  ">
                            <tr>
                                <th scope="col" class="px-6 py-3">
                                    Product name
                                </th>
                                <!-- <th scope="col" class="px-6 py-3">
                                    Color
                                </th> -->
                                <th scope="col" class="px-6 py-3">
                                    Category
                                </th>
                                <th scope="col" class="px-6 py-3">
                                    Price
                                </th>
                                <th scope="col" class="px-6 py-3">
                                    <span class="sr-only">Edit</span>
                                </th>
                            </tr>
                        </thead>
                        <tbody>
                            {#each latestProduct as product(product.ID)}
                                <tr class="bg-white border-b ">
                                    <th scope="row" class="px-6 py-4 font-medium text-gray-900 whitespace-nowrap">
                                        {product.name}
                                    </th>
                                    <!-- <td class="px-6 py-4">
                                        Silver
                                    </td> -->
                                    <td class="px-6 py-4">
                                        {product.category}
                                    </td>
                                    <td class="px-6 py-4">
                                        ${product.price}
                                    </td>
                                    <td class="px-6 py-4 text-right">
                                        <a href={`/update/${product.ID}`} class="font-medium text-blue-600 dark:text-blue-500 hover:underline">Update</a>
                                    </td>
                                </tr>                               
                            {/each}                           
                        </tbody>
                    </table>
                </div>            
            </div>
        </div>


    </div>
</div>