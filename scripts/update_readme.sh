lead='^<pre>$'
tail='^<\/pre>$'
sed -i -e "/$lead/,/$tail/{ /$lead/{p; r $1
        }; /$tail/p; d }" README.md
