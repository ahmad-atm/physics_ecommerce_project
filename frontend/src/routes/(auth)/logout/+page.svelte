<script>
    import {goto} from "$app/navigation"
    import {onMount} from "svelte"

    onMount(async () => {
      localStorage.removeItem("hospital_token")
      goto("/")
    })
</script>