cd "$(dirname "$0")"

sleep 3s

rm "$1.bak"
mv "$1" "$1.bak"
mv "InstallLib*" $1
rm "$1.bak"
