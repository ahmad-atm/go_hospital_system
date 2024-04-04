<script>
  import PatNav from "../../../../components/patNav.svelte";
  import { onMount } from "svelte";

  let tests = []
  let doctors = []

  onMount(async()=>{
    await fetch("http://localhost:8000/patient_tests", {
        headers:{
            "content-type":"application/json",
            Authorization:`Bearer ${localStorage.getItem('hospital_token')}`
        }
    })
    .then(res=>res.json())
    .then(data => {tests=data; console.log(data)})
    .catch(err=>console.log(err))
  })

</script>

<PatNav />
<div>
    <p>My tests</p>
    {#each tests as test(test.ID)}
        <p>{test.title}</p>
    {/each}
</div>