<script>
  import { onMount } from "svelte";
  import Nav from "../../../components/nav.svelte";

  let carts = []
  onMount(()=>{
        fetch("http://localhost:8000/all_cart", {
            headers:{
                "content-type":"application/json",
                authorization:`Bearer ${localStorage.getItem("ecommerce_token")}`
            },
        })
        .then(res=>res.json())
        .then(data=>{carts=data; console.log(data)})
        .catch(err=>console.log(err))
    })

    function removeCart(id){
        console.log(id)
        fetch("http://localhost:8000/delete_cart", {
            method:"Post",
            headers:{
                "content-type":"application/json",
                authorization:`Bearer ${localStorage.getItem('ecommerce_token')}`
            },
            body:JSON.stringify({
                ID:id
            })
        })
        .then(res=>res.json())
        .then(data=>{console.log(data)})
        .catch(err=>console.log(err))
    }
</script>
<div class="bg-white h-screen p-1">
    <div class="bg-yellow-200 h-screen mx-[0.02rem] rounded-t-2xl mt-[0.02rem] p-5">
        <Nav />

        <div class="w-full shadow-md rounded-3xl">
            <div class="relative overflow-x-auto shadow-lg rounded-3xl">
                    <caption class="p-5 flex items-center justify-center text-md font-semibold text-left rtl:text-right text-gray-900">
                        Shopping Cart
                    </caption>                
                <table class="w-full text-sm text-left rtl:text-right text-gray-500 ">

                    <thead class="text-xs text-gray-700 uppercase   ">
                        <tr>
                            <th scope="col" class="px-6 py-2">
                                Image
                            </th>
                            <th scope="col" class="px-6 py-2">
                                Product Name
                            </th>
                            <th scope="col" class="px-6 py-2">
                                Price
                            </th>
                            <th scope="col" class="px-6 py-2">
                                Edit
                            </th>
                        </tr>
                    </thead>
                    <tbody>
                        {#each carts as cart(cart.ID)}
                            <tr class=" border-b ">
                                <th scope="row" class="px-6 py-4 font-medium text-gray-900 whitespace-nowrap">
                                    <img src={`http://localhost:8000/${cart.productImage}`}  alt="" srcset="" class="object-fill h-32 w-32 rounded-full">
                                </th>
                                <td class="px-6 py-4">
                                    {cart.productName}
                                </td>
                                <td class="px-6 py-4">
                                    NGN {cart.productPrice}
                                </td>
                                <td class="px-6 py-4 text-right flex flex-col space-y-4">
                                    <a href={`/cart/${cart.ID}`} class="font-medium text-blue-600 hover:underline flex items-center justify-center">Confirm</a>
                                    <button on:click={()=> removeCart(cart.ID)} class="font-medium text-red-600 hover:underline flex items-center justify-center">Remove</button>
                                </td>
                            </tr>                               
                        {/each}                           
                    </tbody>
                </table>
            </div>            
        </div>
    </div>
</div>