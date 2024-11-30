package com.emarket.productservice.controller;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;


@RestController
@RequestMapping("/products")
public class ProductController {

    Logger logger = LoggerFactory.getLogger(ProductController.class);
    
    @GetMapping("/{id}")
    ResponseEntity<String> getProduct(@PathVariable("id") String id){

        logger.info("Received product id: {}", id);
        return ResponseEntity.ok("12345");
    }
}
