# Run script.sh

# {ID: 1, Stock: 12, Name: "shirt", Color: "blue", Code: "AC123", Price: 12.4, Publish: true, Date: "December"},

curl "http://127.0.0.1:8000/filter-product?id=1&stock=12&name=shirt&color=blue&code=AC123&price=12.4&publish=true&date=December" \
--request GET \
-vvv ; echo "" | cat -e

