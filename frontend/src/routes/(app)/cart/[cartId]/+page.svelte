<script>
// @ts-nocheck

  import Nav from "../../../../components/nav.svelte";
  import { onMount} from "svelte";
  import { page } from "$app/stores";
  import { goto } from "$app/navigation";
  const id = $page.params.cartId;
  let cart = {}
  let name;
  let price;
  let quantity;
  let totalPrice = Number(price) * Number(quantity);
  let additionalInfo = ""
  let x = 1;

  onMount(()=>{
    fetch(`http://localhost:8000/get_cart/${id}`,{
      headers:{
                "content-type":"application/json",
                authorization:`Bearer ${localStorage.getItem('ecommerce_token')}`
            }
        })
        .then(res=>res.json())
        .then(data=>{
          cart=data
          name = cart.productName;
          price = Number(cart.productPrice);
          quantity = Number(cart.quantity);
          totalPrice = price * quantity;
        })
        .catch(err=>console.log(err))
  })

  function confirmOrder(id){
        fetch("http://localhost:8000/confirm_cart", {
            method:"Post",
            headers:{
                "content-type":"application/json",
                authorization:`Bearer ${localStorage.getItem('ecommerce_token')}`
            },
            body:JSON.stringify({
                ID:id,     
                totalPrice: (totalPrice + 1000).toString(), 
                status:"Order_Confirmed",
                quantity: quantity.toString(),
                additionalInfo:additionalInfo
            })
        })
        .then(res=>res.json())
        .then(data=> goto('/shop'))
        .catch(err=>console.log(err))
    }
</script>


<div class="bg-white h-screen p-1">
    <div class="bg-yellow-200 h-screen mx-[0.02rem] rounded-t-2xl mt-[0.02rem] p-5">
        <Nav />
        <div class="flex flex-col items-center mt-24">
            <div class="flex gap-x-3">
                <img src={`http://localhost:8000/${cart.productImage}`}  alt="" srcset="" class="object-cover h-64 w-64 rounded-lg">
                <div class="pt-5 flex flex-col gap-y-3 px-"> 
                    <p class="font-medium text-gray-800">Product Name: {name}</p>
                    <p class="font-medium text-gray-800">Product Price: NGN {price}</p>
                    <div class="flex items-center space-x-2">
                        <button class="bg-gray-100 p-3 rounded-md hover:bg-gray-300 font-semibold" 
                            on:click={()=> {quantity--; 
                                if(quantity < 0){
                                    quantity=0;                                    
                                }totalPrice = price * quantity;
                                }}>-</button>
                            <p>{quantity}</p>
                        <button class="bg-gray-100 p-3 rounded-md hover:bg-gray-300 font-semibold" 
                            on:click={()=> {
                                quantity++;
                                totalPrice = price * quantity;
                            }}>+</button>
                    </div>
                    <p class="font-medium text-gray-800">Delivery Price: NGN 1000</p>
                    <input type="text" bind:value={additionalInfo} placeholder="Additional Information" 
                        class="p-2 w-[30rem] h-[5rem] rounded-lg">
                </div>
            </div>
 
            <button on:click={()=> confirmOrder(cart.ID, cart.name, totalPrice)} 
                class="px-2 py-5 mt-5 bg-black rounded-3xl text-yellow-300 font-bold w-full hover:bg-gray-800">Confirm Order - NGN {totalPrice + 1000}</button>                     
        </div>
    </div>
</div>