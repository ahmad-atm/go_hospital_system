<script>
    import { onMount } from "svelte";
  import LabNav from "../../../../components/labNav.svelte";
    
  let profile = {}
    onMount(()=>{
        fetch("http://localhost:8000/lab_profile",{
            headers:{
              "Content-Type":"application/json",
              Authorization:`Bearer ${localStorage.getItem("hospital_token")}`
            }
        })
        .then(res => res.json())
        .then(data=> {console.log(data); profile=data})
        .catch(err => console.log(err))
    })
</script>

<div>
    <LabNav />
    <p>Profile</p>
</div>