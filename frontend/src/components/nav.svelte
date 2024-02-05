<script>
    import Icon from "@iconify/svelte";
</script>

<nav class="bg-none flex justify-between items-center px-10 py-1 border-b border-orange-300">
    <div class="flex space-x-3 items-center">
        <a href="/home" class="bg-yellow-100 rounded-full p-3">                
            <Icon icon="tabler:shopping-bag"  class="text-3xl text-orange-600"/>                
        </a>    
        <a href="/home" class="rounded-full p-2 ">
            <Icon icon="ion:menu" class="text-2xl text-orange-600"/>
        </a>            
    </div>

    <div class="flex items-center space-x-3">
        <a href="/home" class="font-bold text-sm text-orange-600 hover:text-orange-600">Home</a>
        <a href="/shop" class="font-bold text-sm hover:text-orange-600">Shop</a>
        <a href="/contact" class="font-bold text-sm hover:text-orange-600">Contact</a>
        <a href="/contact" class="font-bold text-sm hover:text-orange-600">About Us</a>
    </div>
    <div class="flex items-center space-x-4">
        <a href="/cart" class="bg-yellow-400 hover:bg-yellow-500 rounded-full p-2 text-orange-600 hover:text-yellow-200">
            <Icon icon="fluent:cart-24-filled"  class="text-xl"/>
        </a>
        <a href="/profile" class="bg-yellow-400 hover:bg-yellow-500 rounded-full p-2 text-orange-600 hover:text-yellow-200">
            <Icon icon="bi:person-fill" class="text-xl"/>
        </a>
        <a href="/logout" class="bg-yellow-400 hover:bg-yellow-500 rounded-full p-2 text-orange-600 hover:text-yellow-200">
            <Icon icon="fe:logout" class="text-xl"/>
        </a>                
    </div>
</nav>