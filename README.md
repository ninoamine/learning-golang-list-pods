# learning-golang-list-pods
for learning golan



This small script use the flag lib to get the namespace and the kubeconfig
the most intersting thing is that it returns a pointer

Than we need to init a kubernetes client by using the right conf

we used flag so we can use BuildConfigFromFlags to get the kubeconfig
than we can init a client with the config