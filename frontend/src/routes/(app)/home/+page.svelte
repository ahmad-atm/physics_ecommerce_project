<script>
  import Nav from "../../../components/nav.svelte";

</script>
<div class="bg-white h-screen p-1">
    <div class="bg-yellow-200 mx-[0.02rem] rounded-t-2xl mt-[0.02rem] p-5">
        <Nav />
        <div class="p-2 mt-10 mx-16 h-80">
           <img src="bf.jpg"  alt="" class="h-full w-full rounded-3xl object-cover">
        </div>

        <div class="pt-7">
            <p class="text-lg font-bold text-gray-700 flex items-center justify-center py-3 border-b-2 border-orange-300">Products</p>
            <div class="flex flex-wrap justify-center items-center space-x-10 gap-y-10 pt-10">
                <div class="bg-white h-60 w-60 p-3 flex items-center justify-center rounded-2xl">
                    <img src="mac.jpg"  alt="" class="h-full w-full rounded-3xl object-contain">
                </div>
                <div class="bg-white h-60 w-60 p-3 flex items-center justify-center rounded-2xl">
                    <img src="iphone.jpg"  alt="" class="h-full w-full rounded-3xl object-contain">
                </div>
                <div class="bg-white h-60 w-60 p-3 flex items-center justify-center rounded-2xl">
                    <img src="console.jpg"  alt="" class="h-full w-full rounded-3xl object-contain">
                </div>
                <div class="bg-white h-60 w-60 p-3 flex items-center justify-center rounded-2xl">
                    <img src="accessories.jpg"  alt="" class="h-full w-full rounded-3xl object-contain">
                </div>
                <div class="bg-white h-60 w-60 p-3 flex items-center justify-center rounded-2xl">
                    <img src="wallpaper.jpg"  alt="" class="h-full w-full rounded-3xl object-contain">
                </div>
                <div class="bg-white h-60 w-60 p-3 flex items-center justify-center rounded-2xl">
                    <img src="iwatch.jpg"  alt="" class="h-full w-full rounded-3xl object-contain">
                </div>
                <div class="bg-white h-60 w-60 p-3 flex items-center justify-center rounded-2xl">
                    <p>item</p>
                </div>
                <div class="bg-white h-60 w-60 p-3 flex items-center justify-center rounded-2xl">
                    <p>item</p>
                </div>                             
            </div>
        </div>
    </div>
</div>