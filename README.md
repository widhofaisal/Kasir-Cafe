# Kasir Cafe by Widho Faisal Hakim

Endpoint :
1. GET      Hello_world         ✅      /hello              admin_controller.go 
2. POST     Login               ✅      /admins/login       admin_controller.go 
3. GET      Get_admins          ✅      /admins             admin_controller.go 
4. POST     Add_product         ✅      /products           product_controller.go 
5. GET      Get_products        ✅      /products           product_controller.go 
6. GET      Get_product_by_id   ✅      /products:id        product_controller.go 
7. PUT      Update_product      ✅      /products:id        product_controller.go 
8. DELETE   Delete_product      ✅      /products:id        product_controller.go 
9. POST     Add_cart            ✅      /carts              cart_controller.go   
10. PUT     Update_cart         ✅      /carts              carts_controllers.go  
11. GET     Get_carts           ✅      /carts              carts_controllers.go  
12. GET     Get_cart_by_id      ✅      /carts/:id          carts_controllers.go  
13. DELETE  Delete_cart         ✅      /carts/:id          carts_controllers.go  
14. GET     Get_nota            ✅      /carts/nota         carts_controllers.go  
15. POST    Add_payment         ✅      /payments           payment_controllers.go  
16. PUT     Paid                ✅      /payments/:id       payment_controllers.go  
17. GET     Total_income        ✅      /income             payment_controllers.go  



Penilaian Mini Project :
1. MVP :
  - Login                               ✅
  - CRUD admin                          ✅
  - CRUD payment                        ✅ 
  - CRUD product                        ✅    
  - Midtrans                            ❌
2. Git workflow                         ✅ 
3. Presentation                         ❌
4. Tech innovation                      ❌
5. Design ERD (figma)                   ❌
6. Design Rest API (swagger/postman)    ❌
7. Backend project :
   - ORM                                ✅
   - MVC                                ✅
   - Unit testing                       ❌
   - Middleware JWT                     ✅
   - 3rd party API (midtrans)           ❌
8. Deployment :     
   - Docker                             ✅
   - Compute service (AWS EC2 & RDS)    ✅
   - CI/CD                              ❌

fix :
- enkripsi password admin, tembak db input hashing nya, berarti di endpoint 2 tambahi encode password (text ke hash)
  
to do :
- bikin table payment dan lakukan payment manual dulu
- implementasi midtrans