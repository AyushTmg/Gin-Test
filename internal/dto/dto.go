package dto

// ! Validation in Gin DTOs
/*
Gin provides a powerful validation mechanism that allows you to validate incoming request
data using struct tags. You can specify validation rules for each field in your DTO
(Data Transfer Object) struct, and Gin will automatically validate the incoming data against
those rules.

1- Required
binding:"required" - This tag indicates that the field is required

2- Email
binding:"email" - This tag indicates that the field must be a valid email address

3- Min and Max Length
binding:"min=6" - This tag indicates that the field must have a minimum length of 6 characters
binding:"max=20" - This tag indicates that the field must have a maximum length of 20 characters

4- len
binding:"len=10" - This tag indicates that the field must have an exact length of 10 characters

5- OneOf
binding:"oneof=admin user guest" - This tag indicates that the field must be one of the specified values (admin, user, or guest)

6- Custom Validation (not included in the built-in validators)
You can also create custom validation functions and use them in your struct tags. For example:
binding:"customValidator" - This tag indicates that the field should be validated using a custom validation function named customValidator

7- Url
binding:"url" - This tag indicates that the field must be a valid URL

8- UUID
binding:"uuid" - This tag indicates that the field must be a valid UUID


Gin will automatically handle the validation based on the tags you specify in your DTO struct. If the validation fails, Gin will return an error response with details about the validation failure.
*/

type CreateUserRequest struct {
	FullName string `json:"full_name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8,max=16"`
}

type UserLoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}
