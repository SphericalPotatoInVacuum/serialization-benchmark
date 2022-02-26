lead='^<!--results begin-->$'
tail='^<!--results finish-->$'
sed -i -e "/$lead/,/$tail/{ /$lead/{p; r $1
        }; /$tail/p; d }" README.md
sed -i -e "s/-|/:|/g; s/ | /\`|\`/g; s/| /|\`/g; s/ |/\`|/g" README.md
