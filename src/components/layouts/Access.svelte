<script>
  import { Router, Route } from 'svelte-routing';
  import SignIn from '../pages/access/SignIn.svelte';
  import SignUp from '../pages/access/SignUp.svelte';
  import ResetPassword from '../pages/access/ResetPassword.svelte';
  export let basepath = '/';

</script>
  
<style>
</style>
  
<Router basepath="{basepath}">
  <Route path="/sign-in" component={SignIn} />
  <Route path="/sign-up" component={SignUp} />
  <Route path="/reset-password" component={ResetPassword} />
</Router>