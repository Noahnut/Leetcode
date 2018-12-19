#include <iostream>
#include <string>
#include <vector>
#include "add2num.h"


using namespace std;

void
add2number(string number_one, string number_two)
{
    add2number_class add2number_class;
    add2number_class.link_list_create(number_one, number_two);   
}

add2number_class::~add2number_class()
{
   node *current;
   for(size_t i = 0; i < node_header.size(); i++){
      current = node_header[i];
      while(current != 0){
      node *temp = current->next;
      delete current;
      current = temp;
      }
   }
}


void
add2number_class::link_list_create(const string &number_one,const string &number_two)
{
   node_header.push_back(block_generator(number_one));
   node_header.push_back(block_generator(number_two));
   sum_the_two_number();
}

void 
add2number_class::sum_the_two_number()
{
   node *current_one; node *current_two; node *current; node *header;
   current_one = node_header[0]; current_two = node_header[1];
   int current_number = current_one->number + current_two->number;
   header = header_generator(current_number);
   current = header;

   current_one = current_one->next; current_two = current_two->next;

   for(auto &i : node_header){
       while(current_one != 0 || current_two != 0){
          node *new_node = new node;
          make_sum(current_one, current_two, new_node, current->up_number); 
          current->next = new_node;
          current = new_node;
          update_current(current_one, current_two,current);
          if(current->up_number == true && current_one == 0 && current_two == 0){
             node *next_node = new node;
             next_node->number = 1;
             current->next = next_node;
             current = next_node;
          }
       }
   }
   node_header.push_back(header);
   print_the_link_list(node_header[2]);
}

void
add2number_class::make_sum(node *&current_one, node *&current_two, node *&new_node, bool up)
{
   int current_number;
   if(current_one != 0 && current_two != 0){
     current_number = current_one->number + current_two->number;
   }
   else if(current_one != 0 && current_two == 0){current_number = current_one->number;}
   else if(current_one == 0 && current_two != 0){current_number = current_two->number;}
   if(current_number >= 10){
      new_node->up_number = true;
      current_number = current_number - 10;
   }
   if(up == true){current_number = current_number + 1;}       
   new_node->number = current_number;
}

void 
add2number_class::update_current(node *&current_one, node *&current_two, node *&current)
{
   if(current_one != 0 && current_two != 0){
      current_one = current_one->next;
      current_two = current_two->next;
   }
   else if(current_one == 0 && current_two != 0){
      node *new_node = new node;
      current_two = current_two->next;
   }
   else if(current_one != 0 && current_two == 0){
      current_one = current_one->next;
   }
}

node*
add2number_class::block_generator(const string &number)
{
   node *head = new node;
   head->number = number[number.size()-1] - '0';
   head->next = head;
   head->head_flag = true;
   node *current;
   current = head;
   for(size_t i = number.size()-2; i >= 0; i--){
      node *new_node = new node;
      new_node->number = number[i] - '0';
      current->next = new_node;
      current = new_node;
      if(i==0){break;}
   }
   return head;
}

node*
add2number_class::header_generator(int &number)
{
   node *header = new node;
   header->head_flag = true;
   if(number >= 10){
      header->up_number = true;
      number = number - 10;
   }
   header->number = number;
   return header;
}

void 
add2number_class::print_the_link_list(node *headeer_node)
{
   node *current;
   current = headeer_node;
   while(current != 0){
      cout << current->number;
      current = current->next;
   } 
   cout << endl;
}