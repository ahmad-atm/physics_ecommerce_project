<script>
    import {goto} from "$app/navigation"

    let name = "";
    let email = "";
    let password = "";
    let role = "Customer"
      
    function register(){
        fetch("http://localhost:8000/register", {
            method:"POST",
            headers:({"Content-Type":"application/json"}),
            body:JSON.stringify({ name,
                email,
                password,
                role })
        })
        .then(res=> res.json())
        .then(data=> {
            console.log(data)            
            goto("/login")
        })
        .catch(err => console.log(err))
    }     
</script>

<div class="flex h-screen items-center justify-center bg-gray-50">
    <!-- <div class="w-[50%] flex  items-center justify-center bg-blue-600 border-4 rounded-lg border-white">
        <p>Ecommerce Register</p>
    </div> -->

    <div class="flex items-center justify-center">
        <form on:submit|preventDefault={register}>
            <div class="gap-y-7 flex flex-col justify-center bg-white shadow-lg shadow-gray-300 p-10 rounded-2xl">
                <p class="flex justify-center items-center font-medium text-lg text-gray-700">Register</p>            
                <div class="flex flex-col gap-y-5">
                    <div class="flex space-x-4 items-center px-5">
                        <label for="name" class="text-gray-700">Name:</label>
                        <input type="name" name="name" id="name" bind:value={name} 
                            placeholder="Name" class="p-2 mb-2 border-b-[0.17rem] border-gray-400 outline-none focus:border-blue-600"
                            autocomplete="name" required />                                
                    </div>
                    <div class="flex space-x-4 items-center px-5">
                        <label for="email" class="text-gray-700">Email:</label>
                        <input type="email" name="email" id="email" bind:value={email} 
                            placeholder="Email" class="p-2 mb-2 border-b-[0.17rem] border-gray-400 outline-none focus:border-blue-600"
                            autocomplete="name" required />                                
                    </div>
                    <div class="flex space-x-4 items-center">
                        <label for="password" class="text-gray-700">Password:</label>
                        <input type="password" name="password" id="password" 
                            autocomplete="current-password" bind:value={password} 
                            placeholder="Password" class="p-2 mb-2 border-b-[0.17rem] border-gray-400 outline-none focus:border-blue-600" 
                            required/>
                    </div>
                </div>
                <button type="submit" class="bg-blue-600 p-3 rounded-lg text-gray-200 hover:bg-blue-400">Register</button>
                <a href="/login" class="flex items-center justify-center text-blue-600">Login</a>
            </div>
        </form>        
    </div>
</div>


<style>
    
</style>