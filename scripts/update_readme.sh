lead='^<!--results begin-->$'
tail='^<!--results finish-->$'
sed -i -e "/$lead/,/$tail/{ /$lead/{p; r $1
        }; /$tail/p; d }"  README.md
