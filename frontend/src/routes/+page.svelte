<script>
    import {onMount} from "svelte"
    import {goto} from "$app/navigation";
  
    onMount(() => {
      if(localStorage.getItem("hospital_token") == null){
        goto("/login")
        return
      }else{
        goto("/home")
      }
    })
</script>