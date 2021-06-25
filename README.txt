Steps to build the client code

1. Set the repo as the workspace root

    cd k8s-codegen-basic
    WSROOT=$(pwd)

2. Make the directory structure

    github.com
        └── user
            └── repo
                ├── go.mod
                ├── go.sum
                └── pkg
                    └── apis
                        └── k8s.crd.io
                            └── v1alpha1
                                ├── doc.go
                                ├── register.go
                                └── types.go

    directory structure mimics the module name github.com/user/repo

3. Go to the repo directory and initialize the module

    cd github.com/user/repo
    MODULE=$(pwd)
    go mod init github.com/user/repo

4. Few things to note

    a. k8s.crd.io is the group name.  it is mentined in doc.go "+groupName=k8s.crd.io" and in register.go scheme group
    b. "+genclient" is added to the struct definition in types.go without empty line between comment and struct definition

5. Setup the Generator path

    GEN = <path to code-generator>

6. Execute the generator

    cd $WSROOT/$MODULE // i.e inside the module directory
    $GEN/generate-groups.sh all $MODULE/pkg/client $MODULE/pkg/apis "k8s.crd.io:v1alpha1" -h $GEN/hack/boilerplate.go.txt -o $WSROOT --go-header-file $GEN/hack/boilerplate.go.txt

    This would generate clientset, informers, listers and zz_generated.deepcopy.go
