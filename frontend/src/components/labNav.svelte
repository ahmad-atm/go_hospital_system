<script>
    import { onMount } from "svelte";

    let labProfile = {}

    onMount(()=>{
        fetch("http://localhost:8000/lab_home",{
            headers:{
                "content-type": "application/json",
                Authorization: `Bearer ${localStorage.getItem("hospital_token")}`
            }
        })
        .then(res => res.json())
        .then(data => labProfile=data)
        .catch(err => console.log(err))
    })
</script>

<nav class="flex justify-between space-x-4 border-b-2 py-6 px-14 mb-4 bg-white">
    <div class="flex gap-x-4 items-center">
        <p class="text-3xl text-rose-600 font-extrabold flex items-center justify-center px-7">ICON</p>
        <a href="/lab_home" class="text-rose-600 hover:underline">Home</a>        
        {#if labProfile.firstName == ""}
            <a href="/lab_profile_register" class="text-rose-600 hover:underline">Create Profile</a>
        {:else if labProfile.firstName != ""}
            <a href="/doctor_list" class="text-rose-600 hover:underline">Doctors</a>   
            <a href="/patient_list" class="text-rose-600 hover:underline">Patients</a>   
            <a href="/tests" class="text-rose-600 hover:underline">Tests</a>   
        {/if}        
    
    </div>
    <div class="flex gap-x-4 items-center">
        {#if labProfile.firstName != ""}
            <a href="/lab_profile" class="text-rose-600 hover:underline">Profile</a>
        {/if} 
        <a href="/logout" class="text-rose-600 hover:underline">Logout</a>
</div>    
</nav>