<script>
  import DocNav from "../../../../components/docNav.svelte";
    import { onMount } from "svelte";
  
    let tests = [];
  
    onMount(()=>{
      fetch("http://localhost:8000/doctor_tests", {
        headers:{
          "Content-type":"application/json",
          Authorization: `Bearer ${localStorage.getItem("hospital_token")}`
        }
      })
      .then(res=>res.json())
      .then(data=>{tests=data; console.log(tests)})
      .catch(err=>console.log(err))
    })
  </script>
  
  <DocNav />
<div class="p-8">
  <div class="flex flex-wrap space-x-10 gap-y-10 pt-5">
    {#each tests as test(test.ID)}
      <a href={`/doctor_tests/${test.ID}`} class="bg-gray-100 h-72 w-48 rounded-2xl flex flex-col space-y-2 items-center  shadow-xl">
          <div class="flex h-52 bg-white w-full rounded-2xl">
              <img src={`http://localhost:8000/${test.patientImage}`}  alt="" srcset="" class="object-cover h-full w-full rounded-t-2xl">
          </div>     
          <div class="flex flex-col items-center">
            <p class="text-gray-700 font-semibold text-sm">{test.patientFirstName} {test.patientLastName}</p>
            <p class="text-green-900 font-normal text-sm">{test.title}</p>
            <p class="text-green-900 font-normal text-xs">{test.status}</p>
          </div>               
      </a>  
      {:else}
        <div class="flex justify-center items-center">
          <p class="text-gray-700 font-medium flex items-center justify-center">No Patient</p>         
        </div>     
    {/each}                              
  </div>
</div>