<script>
  import LabNav from "../../../../components/labNav.svelte";
  import { onMount } from "svelte";
  import {goto} from "$app/navigation"

  let tests = [];

  onMount(()=>{
    fetch("http://localhost:8000/all_tests", {
      headers:{
        "Content-type":"application/json",
        Authorization: `Bearer ${localStorage.getItem("hospital_token")}`
      }
    })
    .then(res=>res.json())
    .then(data=>tests=data)
    .catch(err=>console.log(err))
  })

  function deleteTest(testId){
        fetch(`http://localhost:8000/delete_test`,{
            method:"POST",
            headers:{
                  "content-type": "application/json",
                Authorization: `Bearer ${localStorage.getItem("hospital_token")}`              
            },
            body:JSON.stringify({
                ID:testId,               
            })
      })
      .then(res => res.json())
      .then(data => goto("/tests"))
      .catch(err => console.log(err))
    }
</script>

<LabNav />
<div class="p-8">
  <div class="flex space-x-10 ml-40">
      <a href="/test_create" class="text-white hover:bg-rose-400 bg-rose-500 p-3 rounded-lg">Add Test</a>
  </div>

  <div class="px-10">
    {#each tests as test(test.ID) }
    <div class="mx-24 mt-3 py-7 p-5 border-2 border-rose-400 rounded-lg w-[50%] flex justify-between">
        <p  class="text-lg text-gray-800">{test.title}</p>
        <div class="flex flex-col space-y-3">
          <a href={`/test_update/${test.ID}`} class="px-8 py-3 rounded-xl bg-blue-600 hover:bg-blue-400 text-white">Update</a>
          <a on:click={(testId)=> deleteTest(test.ID)} href="/tests" class="px-8 py-3 rounded-xl bg-rose-600 hover:bg-rose-400 text-white">Delete</a>
        </div>
    </div>
    {/each} 
  </div>
</div>