// Auto-generated from {{name}} ... DO NOT EDIT

include "baseplate.thrift"

{{#service}}
  service {{name}} extends baseplate.BaseplateService {
    bool is_healthy();
    {{#method}}
      {{^health_check}}
      {{outputType}} {{name}} (1: {{inputType}} request);
      {{/health_check}}
    {{/method}}
  }
{{/service}}

{{#messageType}}
  {{thrift_type}} {{name}} {
      {{#field}}
        {{number}}: {{&thrift_type}} {{name}};
      {{/field}}
      {{#oneofs}}
        999: {{name}}{{key}} {{key}}Selection;
      {{/oneofs}}
  }
  
  {{#oneofs}}
    enum {{name}}{{key}} {
      {{#value}}
        {{inner_name}},
      {{/value}}
    }
  {{/oneofs}}
{{/messageType}}