<script>
    import { onMount } from "svelte";
    import {Icon} from "@iconify/svelte"

    let doctor_profile = {}
    onMount(()=>{
        fetch("http://localhost:8000/doctor_home",{
            headers:{
                "content-type": "application/json",
                Authorization: `Bearer ${localStorage.getItem("hospital_token")}`
            }
        })
        .then(res => res.json())
        .then(data => doctor_profile=data)
        .catch(err => console.log(err))
    })
</script>

<nav class="flex flex-col h-screen justify-between space-x-4 border-b-2 py-6 px-14 bg-teal-800">
    <div class="flex flex-col gap-y-10 items-center">
        <p class="text-3xl text-white font-extrabold flex items-center justify-center px-7">
            <Icon icon="fxemoji:hospital" />
        </p>
        <a href="/doctor_home" class="text-white hover:underline">Home</a>        
        {#if doctor_profile.firstName == ""}
            <!-- <a href="/doctor_profile_register" class="text-white hover:underline">Create Profile</a> -->
        {:else if doctor_profile.firstName != ""}
            <a href="/appointments" class="text-white hover:underline">
                <Icon icon="teenyicons:appointments-solid" width=30/>
            </a>   
            <!-- <a href="/doctor_tests" class="text-white hover:underline">Patients</a>    -->
        {/if}        
    
    </div>
    <div class="flex flex-col gap-y-4 items-center">
        {#if doctor_profile.firstName != ""}
            <a href="/doctor_profile" class="text-white hover:underline">
                <Icon icon="gg:profile" width=30/>
            </a>
        {/if} 
            <a href="/logout" class="text-white hover:underline">
                <Icon icon="heroicons-outline:logout" width=30 />
            </a>
    </div>    
</nav>