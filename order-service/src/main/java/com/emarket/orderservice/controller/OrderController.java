package com.emarket.orderservice.controller;

import com.emarket.orderservice.models.Order;
import com.emarket.orderservice.services.OrderService;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;


@RestController
@RequestMapping("/orders")
public class OrderController {

    Logger log = LoggerFactory.getLogger(OrderController.class);

    @Autowired
    OrderService orderService;

    @PostMapping
    ResponseEntity<Order> createOrder(@RequestBody Order order){
        log.info("Order created");
        Order order1 = orderService.createOrder(order);
        return ResponseEntity.status(HttpStatus.CREATED).body(order1);
    }

    @GetMapping("/{id}")
    ResponseEntity<Order> getOrder(@PathVariable("id") String id){
        log.info("Order id {}",id);
        return ResponseEntity.status(HttpStatus.OK).body(orderService.getOrder(id));
    }

    @GetMapping("/get/{id}")
    Order getOrders(@PathVariable("id") String id){
        log.info("Order id {}",id);
        return orderService.getOrder(id);
    }
}
