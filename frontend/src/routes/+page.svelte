<script>
  import { onMount } from "svelte";
  import {goto} from "$app/navigation";
  
  onMount(() => {
    if(localStorage.getItem("ecommerce_token") == null){
      goto("/landing")
      return
    }else{
      goto("/home")
    }
  })
</script>

<h1 class="text-3xl font-bold underline">
    Hello world!
</h1>