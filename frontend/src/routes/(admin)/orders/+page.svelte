<script>
    // @ts-nocheck
    
    import Icon from '@iconify/svelte';
    import AdminNav from '../../../components/adminNav.svelte';
    import {onMount} from "svelte"

    let confirmed_orders = []

    onMount(async()=>{
        await fetch("http://localhost:8000/all_confirmed_cart", {
            headers:{
                "content-type":"application/json",
                authorization:`Bearer ${localStorage.getItem('ecommerce_token')}`
            }
        })
        .then(res=>res.json())
        .then(data=>{confirmed_orders=data; console.log(confirmed_orders)})
        .catch(err=>console.log(err))   
    })
</script>
    
    
<div class="w-screen h-screen flex">
    <!-- <Nav /> -->
    <AdminNav />
    <div class="bg-white w-full h-full shadow-md">
        <div class="relative overflow-x-auto shadow-md sm:rounded-lg">
            <table class="w-full text-sm text-left rtl:text-right text-gray-500 ">
                <caption class="p-5 text-md font-semibold text-left rtl:text-right text-gray-900 bg-white ">
                    Latest orders
                </caption>
                <thead class="text-xs text-gray-700 uppercase bg-gray-50  ">
                    <tr>
                        <th scope="col" class="px-6 py-3">
                            Product Name
                        </th>
                        <th scope="col" class="px-6 py-3">
                            Quantity
                        </th>
                        <th scope="col" class="px-6 py-3">
                            Total Price
                        </th>
                        <th scope="col" class="px-6 py-3">
                            Status
                        </th>
                        <th scope="col" class="px-6 py-3">
                            <span class="sr-only">Edit</span>
                        </th>
                    </tr>
                </thead>
                <tbody>
                    {#each confirmed_orders as orders(orders.ID)}
                        <tr class="bg-white border-b ">
                            <th scope="row" class="px-6 py-4 font-medium text-gray-900 whitespace-nowrap">
                                {orders.productName}
                            </th>
                            <!-- <td class="px-6 py-4">
                                Silver
                            </td> -->
                            <td class="px-6 py-4">
                                {orders.quantity}
                            </td>
                            <td class="px-6 py-4">
                                NGN {orders.totalPrice}
                            </td>
                            <td class="px-6 py-4">
                                {orders.status}
                            </td>                            
                            <td class="px-6 py-4 text-right">
                                <a href="/" class="font-medium text-blue-600 dark:text-blue-500 hover:underline">Process</a>
                            </td>
                        </tr>                               
                    {/each}                           
                </tbody>
            </table>
        </div>            
    </div> 
</div>