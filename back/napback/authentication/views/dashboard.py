# from django.shortcuts import render
# from django.contrib.auth.decorators import login_required

# @login_required(login_url='login')
# def dashboard(request):
#     # Examine request
#     print(request.method)  # HTTP method used for the request


from django.shortcuts import render
from django.contrib.auth.decorators import login_required
from authentication.views.nats_publisher import send_nats_message

@login_required(login_url='login')
def dashboard(request):
    sourceNode = type = targetNode = None

    if 'connName' in request.GET:
        # Parameter 'connName' exists, set type to 'CreateInfrastructure'
    
        type = 'CreateInfrastructure'
        
        if (request.GET['sourceNode'] == 'Node 1'):
            sourceNode = 'team1-NE-1'
        elif (request.GET['sourceNode'] == 'Node 2'):
            sourceNode = 'team1-NE-2'
        elif (request.GET['sourceNode'] == 'Node 3'):
                sourceNode = 'team1-NE-3'
                    
        if (request.GET['targetNode'] == 'Node 1'):
            targetNode = 'team1-NE-1'
        elif (request.GET['targetNode'] == 'Node 2'):
            targetNode = 'team1-NE-2'
        elif (request.GET['targetNode'] == 'Node 3'):
            targetNode = 'team1-NE-3'

    elif 'serviceRate' in request.GET:
        # Parameter 'serviceRate' exists, set type to 'CreateService'
        type = 'CreateService'
        
        if (request.GET['sourceNode'] == 'Node 1'):
            sourceNode = 'team1-NE-1'
        elif (request.GET['sourceNode'] == 'Node 2'):
            sourceNode = 'team1-NE-2'
        elif (request.GET['sourceNode'] == 'Node 3'):
                sourceNode = 'team1-NE-3'
                    
        if (request.GET['targetNode'] == 'Node 1'):
            targetNode = 'team1-NE-1'
        elif (request.GET['targetNode'] == 'Node 2'):
            targetNode = 'team1-NE-2'
        elif (request.GET['targetNode'] == 'Node 3'):
            targetNode = 'team1-NE-3'
    else:
        # Parameter 'deleteType' exists, set type to 'DeleteInfrastructure'
        type = 'DeleteInfrastructure'

    if type == 'CreateInfrastructure':
        connectionName = request.GET['connName']
        # Print for debugging
        print(f"CreateInfrastructure Form - Source Node: {sourceNode}, Target Node: {targetNode}, Connection Name: {connectionName}")
        message = f"CreateInfrastructure,{sourceNode},{targetNode},{connectionName}"
        send_nats_message('create.infrastructure', message)

    elif type == 'CreateService':
            serviceRate = request.GET['serviceRate']
            numberOfService = request.GET['numberOfService']
            print(f"CreateService Form - Source Node: {sourceNode}, Target Node: {targetNode}, Service Rate: {serviceRate}, Number of Service: {numberOfService}")
            message = f"CreateService,{sourceNode},{targetNode},{serviceRate},{numberOfService}"
            send_nats_message('create.service', message)

    elif type == 'DeleteInfrastructure':
            deleteType = request.GET.getlist('deleteType')  # In case multiple checkboxes are selected
            infraName = request.GET.getlist("infraName",None)
            serviceName=request.GET.getlist("serviceName",None)
            if(infraName):
                print(f"DeleteInfrastructure Form - Delete Type: {deleteType}, Infra Name: {infraName}","Service Name: {serviceName}")
            # Publish message for deleting infrastructure if deleteType is 'Infrastructure' 
            if 'Infrastructure' in deleteType:
                message = f"DeleteInfrastructure,{infraName}"
                send_nats_message('delete',message)
            # Publish message for deleting service if deleteType is 'Service'
            if 'Service' in deleteType:
                message = f"DeleteService,{serviceName}"
                send_nats_message('delete',message)
            # Publish message for deleting both infrastructure and service if deleteType is 'Both'
            if 'Both' in deleteType:
                message = f"DeleteBoth,{infraName},{serviceName}"
                send_nats_message('delete',message)


    return render(request, 'hpanel/user_panel.html')