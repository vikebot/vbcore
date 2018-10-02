#! /bin/bash
# Set global email and name of commiter
git config --global user.email "vikebot@harwoeck.at"
git config --global user.name "Vikebot CI"

# Get the hash of the commit that triggered this commit
COMMIT_TRIGGER=$(git rev-parse --verify HEAD)

commitDependencyGraph() {
    # Generate new diagram
    goviz -i github.com/vikebot/vbcore -d 1 | dot -Tpng -o diagrams/vbcore-dependency.png
    
    # Commit new diagram
    git add diagrams/vbcore-dependency.png
    git commit -m "Automatic dependency diagram update: vbcore build-$TRAVIS_BUILD_NUMBER"
    git push
}

commitVbdbVersion() {
    # Clone repository we commit to and change to it's directory
    git clone git@github.com:$1/$2.git $2
    cd $2

    # Create version directory if it doesn't exist
    if [ ! -d ".version" ]; then
        mkdir .version
    fi

    # Delete old vbcore version file
    rm -f .version/vbcore.md

    # Create new vbcore version file
    echo "**Automatic dependency update:** vbcore build-$TRAVIS_BUILD_NUMBER https://travis-ci.com/vikebot/vbcore/builds/$TRAVIS_BUILD_ID" >> .version/vbcore.md
    echo "" >> .version/vbcore.md
    echo "**Commit:** _\"$TRAVIS_COMMIT_MESSAGE\"_ https://github.com/vikebot/vbcore/commit/$COMMIT_TRIGGER" >> .version/vbcore.md

    # Commit new file
    git add .version/vbcore.md
    git commit -m "Automatic dependency update: vbcore build-$TRAVIS_BUILD_NUMBER"
    git push

    # Leave directory
    cd ..
}

commitDependencyGraph

commitVbdbVersion "vikebot" "vbdb"
